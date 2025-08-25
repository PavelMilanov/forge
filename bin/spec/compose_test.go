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

func TestComposeInit(t *testing.T) {
	compose := Compose{Tag: "latest"}

	for idx, tmpl := range []string{composeTmpl1, composeTmpl2} {
		tmpfile := fmt.Sprintf("compose-template-%d.yaml", idx)
		if err := os.WriteFile(tmpfile, []byte(tmpl), 0644); err != nil {
			t.Errorf("Ошибка записи файла: %v", err)
		}
		if err := compose.Init(tmpfile, fmt.Sprintf("%d", idx)); err != nil {
			t.Errorf("Ошибка инициализации: %v", err)
		}
		defer os.Remove(tmpfile)
	}
}
