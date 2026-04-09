package env

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/templatevars"
)

type projectState struct {
	Template     string                    `json:"template"`
	Environments map[string]environmentRef `json:"environments"`
}

type environmentRef struct {
	Placement map[string]any `json:"placement"`
	Data      map[string]any `json:"data"`
}

func loadProjectState(ctx context.Context, project string) (projectState, error) {
	secret, err := vault.API.Get(ctx, project)
	if err != nil {
		return projectState{}, err
	}

	state := projectState{
		Environments: map[string]environmentRef{},
	}

	if t, ok := secret.Data["template"].(string); ok {
		state.Template = t
	}

	envsRaw, ok := secret.Data["environments"]
	if !ok {
		return state, nil
	}

	envs, ok := toMap(envsRaw)
	if !ok {
		return state, fmt.Errorf("invalid environments format")
	}

	for envName, envValue := range envs {
		envMap, ok := toMap(envValue)
		if !ok {
			continue
		}
		ref := environmentRef{
			Placement: map[string]any{},
			Data:      map[string]any{},
		}
		if pRaw, ok := envMap["placement"]; ok {
			if p, ok := toMap(pRaw); ok {
				ref.Placement = p
			}
		}
		if dRaw, ok := envMap["data"]; ok {
			if d, ok := toMap(dRaw); ok {
				ref.Data = d
			}
		} else {
			return state, fmt.Errorf("invalid environment format for %q: missing data field", envName)
		}
		state.Environments[envName] = ref
	}

	return state, nil
}

func saveProjectState(ctx context.Context, project string, state projectState) error {
	payload := map[string]any{
		"template":     state.Template,
		"environments": map[string]any{},
	}

	envs := payload["environments"].(map[string]any)
	for name, ref := range state.Environments {
		envs[name] = map[string]any{
			"placement": ref.Placement,
			"data":      ref.Data,
		}
	}

	_, err := vault.API.Put(ctx, project, payload)
	return err
}

func ensureEnvironment(state *projectState, envName string) {
	if state.Environments == nil {
		state.Environments = map[string]environmentRef{}
	}
	ref, ok := state.Environments[envName]
	if !ok {
		ref = environmentRef{
			Placement: map[string]any{},
			Data:      map[string]any{},
		}
	}
	if ref.Placement == nil {
		ref.Placement = map[string]any{}
	}
	if ref.Data == nil {
		ref.Data = map[string]any{}
	}
	setPathIfMissing(ref.Placement, "enabled", false)
	state.Environments[envName] = ref
}

func fillEnvironmentDataFromTemplate(state *projectState, envName string) error {
	if state.Template == "" {
		return fmt.Errorf("template is empty")
	}

	vars, err := extractTemplateVariables(state.Template)
	if err != nil {
		return err
	}

	ensureEnvironment(state, envName)
	ref := state.Environments[envName]
	for key, value := range vars {
		if strings.HasPrefix(key, "placement.") {
			setPathIfMissing(ref.Placement, strings.TrimPrefix(key, "placement."), value)
			continue
		}
		setPathIfMissing(ref.Data, key, value)
	}
	// Единый безопасный дефолт для условного placement в шаблонах swarm.
	setPathIfMissing(ref.Placement, "enabled", false)
	state.Environments[envName] = ref
	return nil
}

func getEnvironment(state projectState, envName string) (environmentRef, bool) {
	ref, ok := state.Environments[envName]
	return ref, ok
}

func parseParam(param string) (string, any, error) {
	parts := strings.SplitN(param, "=", 2)
	if len(parts) != 2 {
		return "", nil, fmt.Errorf("invalid parameter format %q, expected key=value", param)
	}
	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])
	if key == "" {
		return "", nil, fmt.Errorf("empty key in %q", param)
	}
	return key, castValue(value), nil
}

func castValue(raw string) any {
	if v, err := strconv.Atoi(raw); err == nil {
		return v
	}
	if v, err := strconv.ParseBool(raw); err == nil {
		return v
	}
	if strings.Contains(raw, ".") {
		if v, err := strconv.ParseFloat(raw, 64); err == nil {
			return v
		}
	}
	return raw
}

func setPath(root map[string]any, path string, value any) {
	parts := strings.Split(path, ".")
	cur := root
	for i := 0; i < len(parts)-1; i++ {
		key := parts[i]
		next, ok := cur[key]
		if !ok {
			m := map[string]any{}
			cur[key] = m
			cur = m
			continue
		}
		if cast, ok := toMap(next); ok {
			cur = cast
			continue
		}
		m := map[string]any{}
		cur[key] = m
		cur = m
	}
	cur[parts[len(parts)-1]] = value
}

func setPathIfMissing(root map[string]any, path string, value any) {
	parts := strings.Split(path, ".")
	cur := root
	for i := 0; i < len(parts)-1; i++ {
		key := parts[i]
		next, ok := cur[key]
		if !ok {
			m := map[string]any{}
			cur[key] = m
			cur = m
			continue
		}
		cast, ok := toMap(next)
		if !ok {
			m := map[string]any{}
			cur[key] = m
			cur = m
			continue
		}
		cur = cast
	}

	last := parts[len(parts)-1]
	if _, exists := cur[last]; exists {
		return
	}
	cur[last] = value
}

func extractTemplateVariables(templateName string) (map[string]any, error) {
	candidates := []string{
		templateName,
		filepath.Join(config.TEMPLATE_PATH, templateName),
	}

	var (
		content []byte
		err     error
	)
	for _, path := range candidates {
		content, err = os.ReadFile(path)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}

	result, err := templatevars.ExtractFromYAML(content)
	if err != nil {
		return nil, err
	}
	return result.Variables, nil
}

func renderTemplate(templateName string, data map[string]any) (string, error) {
	path := filepath.Join(config.TEMPLATE_PATH, templateName)
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	tmpl, err := template.New(filepath.Base(templateName)).Option("missingkey=error").Parse(string(content))
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func renderContextForEnvironment(ref environmentRef) map[string]any {
	ctx := map[string]any{}
	for key, value := range ref.Data {
		ctx[key] = value
	}
	ctx["placement"] = ref.Placement
	return ctx
}

func marshalPrettyJSON(value any) (string, error) {
	b, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func toMap(value any) (map[string]any, bool) {
	switch typed := value.(type) {
	case map[string]any:
		return typed, true
	default:
		return nil, false
	}
}

func containsSecretNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), "secret not found")
}
