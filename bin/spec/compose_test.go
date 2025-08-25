package spec

import (
	"bytes"
	"html/template"
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

func TestComposeInit(t *testing.T) {
	compose := Compose{Tag: "latest"}

	for _, tmpl := range []string{composeTmpl1, composeTmpl2} {
		tmpl, err := template.New("compose-template").Parse(tmpl)
		if err != nil {
			t.Errorf("Ошибка парсинга шаблона: %v", err)
		}
		var buf bytes.Buffer

		err = tmpl.Execute(&buf, compose)
		if err != nil {
			t.Error(err)
		}

		t.Log(buf.String())
	}
}
