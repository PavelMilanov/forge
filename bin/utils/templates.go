// Package utils реализует вспомогательные функции.
package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/PavelMilanov/forge/config"
)

func GenerateAppConfig(path, name string, tags map[string]string) (string, error) {
	funcMap := template.FuncMap{
		"tag": func(svc string) string {
			if t, ok := tags[svc]; ok {
				return t
			}
			return "undefined"
		},
	}

	tmpl, err := template.New(filepath.Base(path)).
		Funcs(funcMap).
		ParseFiles(path)
	if err != nil {
		return "", err
	}
	fileName := fmt.Sprintf("%s-stack.yml", name)
	output, err := os.Create(filepath.Join(config.CONFIG_PATH, fileName))
	if err != nil {
		return "", err
	}
	defer output.Close()
	if err := tmpl.ExecuteTemplate(output, filepath.Base(path), nil); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	return fileName, nil
}
