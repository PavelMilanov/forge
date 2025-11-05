package spec

import (
	"fmt"
	"os"
	"testing"

	"github.com/PavelMilanov/forge/config"
)

const swarmTmpl1 = `
services:
  mysql:
    image: mysql:latest
    environment:
      - MYSQL_ROOT_PASSWORD=root
      - MYSQL_DATABASE=db
      - MYSQL_USER=root
      - MYSQL_PASSWORD=root
    volumes:
      - mysql-data:/var/lib/mysql

volumes:
	mysql-data:
`

const swarmTmpl2 = `
services:
  mysql:
    image: mysql:{{.Tag}}
    deploy:
      replicas: {{.Replicas}}
    environment:
      - MYSQL_ROOT_PASSWORD=root
      - MYSQL_DATABASE=db
      - MYSQL_USER=root
      - MYSQL_PASSWORD=root
    volumes:
      - mysql-data:/var/lib/mysql

volumes:
	mysql-data:
`

func TestSwarmGenerate(t *testing.T) {
	swarm := Swarm{Tag: "latest", Replicas: 1}
	config.CONFIG_PATH = "."
	for idx, tmpl := range []string{swarmTmpl1, swarmTmpl2} {
		tmpfile := fmt.Sprintf("swarm-template-%d.yaml", idx)
		if err := os.WriteFile(tmpfile, []byte(tmpl), 0644); err != nil {
			t.Fatalf("Ошибка записи файла: %v", err)
		}
		config, err := swarm.Generate(tmpfile, fmt.Sprintf("%d", idx))
		if err != nil {
			t.Fatalf("Ошибка генерации: %v", err)
		}
		defer os.Remove(tmpfile)
		r, err := os.ReadFile(config)
		if err != nil {
			t.Fatalf("Ошибка чтения файла: %v", err)
		}
		t.Log(string(r))
		defer os.Remove(config)
	}
}

func TestSwarmInit(t *testing.T) {
	model := Swarm{}
	model.Init()
	if model.Tag != "latest" || model.Replicas != 1 {
		t.Fatalf("Неверная инициализация модели: %+v", model)
	}
}

func TestSwarmParse(t *testing.T) {
	data := map[string]any{"tag": "test", "replicas": "10"}
	model := Swarm{}
	model.Parse(data)
	if model.Tag != "test" || model.Replicas != 10 {
		t.Fatalf("Неверная инициализация модели: %+v", model)
	}
	t.Logf("%+v", model)
}

func TestSwarmUpdate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		t.Run("1/1 parameter", func(t *testing.T) {
			model := Swarm{}
			model.Init()
			t.Logf("%+v", model)
			input := []string{"tag=test"}
			if error := model.Update(input); error != nil {
				t.Fatalf("Ошибка обновления: %v", error)
			}
			t.Logf("%+v", model)
			if model.Tag != "test" || model.Replicas != 1 {
				t.Fatalf("Неверное обновление: %+v\nВходящие параметры: %s", model, input)
			}
		})
		t.Run("1/2 parameter", func(t *testing.T) {
			model := Swarm{}
			model.Init()
			t.Logf("%+v", model)
			input := []string{"replicas=10"}
			if error := model.Update(input); error != nil {
				t.Fatalf("Ошибка обновления: %v", error)
			}
			t.Logf("%+v", model)
			if model.Tag != "latest" || model.Replicas != 10 {
				t.Fatalf("Неверное обновление: %+v\nВходящие параметры: %s", model, input)
			}
		})
		t.Run("2 parameters", func(t *testing.T) {
			model := Swarm{}
			model.Init()
			t.Logf("%+v", model)
			input := []string{"tag=test", "replicas=10"}
			if error := model.Update(input); error != nil {
				t.Fatalf("Ошибка обновления: %v", error)
			}
			t.Logf("%+v", model)
			if model.Tag != "test" || model.Replicas != 10 {
				t.Fatalf("Неверное обновление: %+v", model)
			}
		})
	})
	t.Run("not valid", func(t *testing.T) {
		input := []string{"tag test"}
		model := Compose{}
		if err := model.Update(input); err != nil {
			t.Log(err)
		}

	})
	t.Run("not valid2", func(t *testing.T) {
		input := []string{"tags=test"}
		model := Compose{}
		if err := model.Update(input); err != nil {
			t.Log(err)
		}

	})
}
