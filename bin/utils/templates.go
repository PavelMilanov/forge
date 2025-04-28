// Package utils реализует вспомогательные функции.
package utils

import (
	"os"
	"text/template"
)

func GenerateAppConfig() {
	tmpl, err := template.ParseFiles("template.yml")
	if err != nil {
		panic(err)
	}
	output, err := os.Create("docker-compose.yaml")
	if err != nil {
		panic(err)
	}
	defer output.Close()

	data := struct {
		Tag string
	}{
		Tag: "3.14",
	}
	err = tmpl.Execute(output, data)
	if err != nil {
		panic(err)
	}
}
