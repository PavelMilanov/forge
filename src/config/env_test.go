package config

import (
	"os"
	"testing"
)

func TestNewEnv(t *testing.T) {
	t.Run("Agent type: portainer", func(t *testing.T) {
		var config = `
vault:
  url: "http://localhost:8080"
  role_id: "xxxx"
  secret_id: "xxxx"

agent:
  type: portainer
  credentials:
    url: "http://localhost:8081"
    key: "xxxx"

`
		if err := os.WriteFile("forge.yml", []byte(config), 0644); err != nil {
			t.Fatal(err)
		}
		env, err := NewEnv(".", "forge.yml")
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", env)
		os.Remove("forge.yml")
	})
	t.Run("Agent type: ssh", func(t *testing.T) {
		var config = `
vault:
  url: "http://localhost:8080"
  role_id: "xxxx"
  secret_id: "xxxx"

agent:
  type: ssh

`
		if err := os.WriteFile("forge.yml", []byte(config), 0644); err != nil {
			t.Fatal(err)
		}
		env, err := NewEnv(".", "forge.yml")
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", env)
		os.Remove("forge.yml")
	})
	t.Run("Agent type: docker", func(t *testing.T) {
		var config = `
vault:
  url: "http://localhost:8080"
  role_id: "xxxx"
  secret_id: "xxxx"

agent:
  type: docker

`
		if err := os.WriteFile("forge.yml", []byte(config), 0644); err != nil {
			t.Fatal(err)
		}
		env, err := NewEnv(".", "forge.yml")
		if err.Error() != "agent type is invalid" {
			t.Error(err)
		}
		t.Logf("%+v", env)
		os.Remove("forge.yml")
	})
}
