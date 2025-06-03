package docker

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewStack(t *testing.T) {
	t.Run("absolute path", func(t *testing.T) {
		curr, err := os.Getwd()
		if err != nil {
			t.Errorf("error: %v", err)
		}
		c, err := NewStack(filepath.Join(curr, "./test/docker-compose.test1.yaml"))
		if err != nil {
			t.Errorf("error: %v", err)
		}
		t.Log(c)
	})
	t.Run("relative path", func(t *testing.T) {
		c, err := NewStack("test/docker-compose.test1.yaml")
		if err != nil {
			t.Errorf("error: %v", err)
		}
		t.Log(c)
	})

}

func TestRegistryLogin(t *testing.T) {
	t.Run("login", func(t *testing.T) {
		err := RegistryLogin("default", "admin", "admin", "http://127.0.0.1:5050")
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}
