package deploy

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

// endpointCompletion возвращает список endpoint-ов для shell completion в командах deploy stack.
func endpointCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if err := loadEndpointAliases(); err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	return endpointAliases, cobra.ShellCompDirectiveNoFileComp
}

// endpointArgs валидирует endpoint-аргумент и ограничивает его списком доступных endpoint-ов.
func endpointArgs(cmd *cobra.Command, args []string) error {
	if err := loadEndpointAliases(); err != nil {
		return err
	}
	cmd.ValidArgs = endpointAliases
	return cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs)(cmd, args)
}

// loadEndpointAliases получает список endpoint-ов из Portainer API и кеширует его
// в переменных endpoints/endpointAliases для валидации аргументов и автодополнения.
func loadEndpointAliases() error {
	cfg, err := api.NewPortainer(
		config.AppConfig.Agent.Credentials.Url,
		config.AppConfig.Agent.Credentials.Key,
		config.AppConfig.Agent.Credentials.Teams)
	if err != nil {
		return err
	}
	data, err := cfg.GetEndpoints()
	if err != nil {
		return err
	}
	endpoints = data
	endpointAliases = endpointAliases[:0]
	for _, item := range endpoints {
		endpointAliases = append(endpointAliases, item.Name)
	}
	return nil
}

// preparePortainerAndEndpoint создает клиент Portainer и резолвит endpoint alias в endpoint ID.
func preparePortainerAndEndpoint(endpointAlias string) (*api.Portainer, int) {
	cfg, err := api.NewPortainer(
		config.AppConfig.Agent.Credentials.Url,
		config.AppConfig.Agent.Credentials.Key,
		config.AppConfig.Agent.Credentials.Teams)
	if err != nil {
		errors.DeployErrors(err)
	}
	endpointID, err := resolveEndpointID(cfg, endpointAlias)
	if err != nil {
		errors.DeployErrors(err)
	}
	return cfg, endpointID
}

// resolveEndpointID ищет endpoint по имени и возвращает его ID.
func resolveEndpointID(cfg *api.Portainer, alias string) (int, error) {
	if alias == "" {
		return 0, fmt.Errorf("endpoint is required")
	}
	endpoint, found, err := cfg.GetEndpointByName(alias)
	if err != nil {
		return 0, err
	}
	if !found {
		return 0, fmt.Errorf("endpoint %q not found", alias)
	}
	return endpoint.ID, nil
}

// buildTemplateContent рендерит custom template Portainer в итоговый stack-файл.
func buildTemplateContent(cfg *api.Portainer, templateName, stackName string) (string, error) {
	tmpl, found, err := cfg.GetTemplateByName(templateName)
	if err != nil {
		return "", err
	}
	if !found {
		return "", fmt.Errorf("template %q not found in Portainer", templateName)
	}
	proj, err := cfg.GetTemplateFile(*tmpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	data := map[string]string{
		"stand": stackName,
	}
	funcMap := template.FuncMap{
		"port": func(s string) string {
			return strings.TrimPrefix(s, "f")
		},
	}
	t, err := template.New("portainer").Option("missingkey=error").Funcs(funcMap).Parse(string(proj.TemplateFile))
	if err != nil {
		return "", err
	}
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
