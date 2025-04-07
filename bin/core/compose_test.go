package core

import (
	"os"
	"testing"
)

var compose Compose
var composeFile = "docker-compose.yaml"

func createFile(name string) {
	text := `services:
  web:
    image: nginx:latest
    container_name: my_nginx
    command: ["nginx", "-g", "daemon off;"]
    ports:
      - "8080:80"
`
	os.WriteFile(name, []byte(text), 0644)
}

func TestCompose(t *testing.T) {
	createFile(composeFile)
	t.Run("run", func(t *testing.T) {
		compose.Command("up", composeFile)
	})
	t.Run("stop", func(t *testing.T) {
		compose.Command("down", composeFile)
	})
	os.Remove(composeFile)
}
