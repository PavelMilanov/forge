package docker

import (
	"context"
	"testing"

	"github.com/docker/docker/api/types/registry"
)

func TestRegistryLogin(t *testing.T) {
	client, err := GetDockerClient()
	if err != nil {
		t.Errorf("Error getting Docker client: %v", err)
	}
	defer client.Close()
	cred := registry.AuthConfig{
		Username:      "admin",
		Password:      "admin",
		ServerAddress: "https://registry.com",
	}
	reg, err := client.RegistryLogin(context.Background(), cred)
	if err != nil {
		t.Errorf("Error logging in to registry: %v", err)
	}
	t.Logf("Logged in to registry: %s", reg.Status)
}

func TestGetContainers(t *testing.T) {
	client, err := GetDockerClient()
	if err != nil {
		t.Errorf("Error getting Docker client: %v", err)
	}
	data, err := GetContainers(client)
	if err != nil {
		t.Errorf("Error getting containers: %v", err)
	}
	t.Log(data)
}
