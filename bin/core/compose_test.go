package core

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewCompose(t *testing.T) {
	t.Run("absolute path", func(t *testing.T) {
		curr, err := os.Getwd()
		if err != nil {
			t.Errorf("error: %v", err)
		}
		c, err := NewCompose(filepath.Join(curr, "./test/docker-compose.yaml"))
		if err != nil {
			t.Errorf("error: %v", err)
		}
		t.Log(c.Dir)
	})
	t.Run("relative path", func(t *testing.T) {
		c, err := NewCompose("test/docker-compose.yaml")
		if err != nil {
			t.Errorf("error: %v", err)
		}
		t.Log(c.Dir)
	})

}
