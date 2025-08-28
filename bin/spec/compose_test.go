package spec

import (
	"fmt"
	"os"
	"testing"
)

const composeTmpl1 = `
services:
  mysql:
    image: mysql:{{.Tag}}
    restart: unless-stopped
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

const composeTmpl2 = `
services:
  redis:
    image: redis:latest
    restart: unless-stopped
    environment:
      - REDIS_PASSWORD=redis
    volumes:
      - redis-data:/data

volumes:
	redis-data:
`

func TestComposeNew(t *testing.T) {
	compose := Compose{Tag: "latest"}

	for idx, tmpl := range []string{composeTmpl1, composeTmpl2} {
		tmpfile := fmt.Sprintf("compose-template-%d.yaml", idx)
		if err := os.WriteFile(tmpfile, []byte(tmpl), 0644); err != nil {
			t.Errorf("Ошибка записи файла: %v", err)
		}
		if err := compose.New(tmpfile, fmt.Sprintf("%d", idx)); err != nil {
			t.Errorf("Ошибка инициализации: %v", err)
		}
		defer os.Remove(tmpfile)
	}
}

func TestComposeInit(t *testing.T) {
	model := Compose{}
	model.Init()
	if model.Tag != "latest" {
		t.Fatalf("Неверная инициализация: %+v", model)
	}
}

func TestComposeParse(t *testing.T) {
	data := map[string]any{"tag": "test"}
	model := Compose{}
	model.Parse(data)
	if model.Tag != "test" {
		t.Fatalf("Неверный парсинг: %s", model.Tag)
	}
	t.Logf("%+v", model)
}

func TestComposeUpdate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		model := Compose{}
		model.Init()
		input := []string{"tag=test"}
		if error := model.Update(input); error != nil {
			t.Fatalf("Ошибка обновления: %v", error)
		}
		if model.Tag != "test" {
			t.Fatalf("Неверное обновление: %s", model.Tag)
		}
	})
	t.Run("not valid", func(t *testing.T) {
		input := []string{"tag test"}
		model := Compose{}
		model.Init()
		if err := model.Update(input); err != nil {
			t.Log(err)
		}
	})
}
