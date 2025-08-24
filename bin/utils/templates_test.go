package utils

import (
	"os"
	"path/filepath"
	"testing"
	"text/template"
)

func TestGenerateAppConfig(t *testing.T) {
	tags := map[string]string{
		"alpine":   "test",
		"nginx":    "test",
		"postgres": "test",
	}
	_, err := GenerateAppConfig("../docker/test/docker-compose.test1.yaml", "test", tags)
	if err != nil {
		t.Error(err)
	}
}

func TestParse(t *testing.T) {
	path := "docker-test.yml"
	fileName := filepath.Base(path + ".generated")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		t.Error(err)
	}
	x := struct {
		Tag      string
		Replicas int
	}{
		Tag:      "test",
		Replicas: 1,
	}
	genFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Error(err)
	}
	defer genFile.Close()
	err = tmpl.Execute(genFile, x)
	if err != nil {
		t.Error(err)
	}
}
