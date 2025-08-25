package spec

import (
	"fmt"
	"os"
	"testing"
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

func TestSwarmNew(t *testing.T) {
	swarm := Swarm{Tag: "latest", Replicas: 1}

	for idx, tmpl := range []string{swarmTmpl1, swarmTmpl2} {
		tmpfile := fmt.Sprintf("swarm-template-%d.yaml", idx)
		if err := os.WriteFile(tmpfile, []byte(tmpl), 0644); err != nil {
			t.Errorf("Ошибка записи файла: %v", err)
		}
		if err := swarm.New(tmpfile, fmt.Sprintf("%d", idx)); err != nil {
			t.Errorf("Ошибка инициализации: %v", err)
		}
		defer os.Remove(tmpfile)
	}

}
