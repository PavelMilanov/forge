package templatevars

import (
	"regexp"
	"strings"

	yaml "go.yaml.in/yaml/v3"
)

var varRegexp = regexp.MustCompile(`\{\{\s*\.([A-Za-z0-9_.-]+)\s*\}\}`)
var standaloneTemplateActionRegexp = regexp.MustCompile(`^\s*\{\{[-]?\s*(.*?)\s*-?\}\}\s*$`)

// VariablesResult содержит переменные, извлеченные из Go template выражений в YAML.
type VariablesResult struct {
	Variables map[string]any `json:"variables"`
}

// ExtractFromYAML извлекает переменные формата {{.path.to.value}} из всех строковых скаляров YAML.
func ExtractFromYAML(content []byte) (VariablesResult, error) {
	normalized := stripStandaloneControlDirectives(string(content))

	var doc any
	if err := yaml.Unmarshal([]byte(normalized), &doc); err != nil {
		return VariablesResult{}, err
	}

	result := VariablesResult{Variables: map[string]any{}}
	walk(doc, func(s string) {
		matches := varRegexp.FindAllStringSubmatch(s, -1)
		for _, match := range matches {
			if len(match) < 2 {
				continue
			}
			key := match[1]
			if _, exists := result.Variables[key]; exists {
				continue
			}
			if isNumericLike(key) {
				result.Variables[key] = 1
			} else {
				result.Variables[key] = ""
			}
		}
	})

	return result, nil
}

func stripStandaloneControlDirectives(content string) string {
	lines := strings.Split(content, "\n")
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		if isStandaloneControlDirectiveLine(line) {
			continue
		}
		out = append(out, line)
	}
	return strings.Join(out, "\n")
}

func isStandaloneControlDirectiveLine(line string) bool {
	m := standaloneTemplateActionRegexp.FindStringSubmatch(line)
	if len(m) < 2 {
		return false
	}

	inner := strings.TrimSpace(m[1])
	if inner == "" {
		return false
	}
	if strings.HasPrefix(inner, "/*") {
		return true
	}

	fields := strings.Fields(inner)
	if len(fields) == 0 {
		return false
	}
	switch fields[0] {
	case "if", "else", "end", "range", "with", "define", "block", "template":
		return true
	default:
		return false
	}
}

func walk(value any, onString func(string)) {
	switch v := value.(type) {
	case map[string]any:
		for _, item := range v {
			walk(item, onString)
		}
	case []any:
		for _, item := range v {
			walk(item, onString)
		}
	case string:
		onString(v)
	}
}

func isNumericLike(key string) bool {
	normalized := strings.ToLower(key)
	return strings.Contains(normalized, "replica") ||
		strings.Contains(normalized, "count") ||
		strings.Contains(normalized, "port") ||
		strings.Contains(normalized, "timeout")
}
