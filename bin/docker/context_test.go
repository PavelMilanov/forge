package docker

import "testing"

func TestGetDockerClient(t *testing.T) {
	client, err := GetDockerClient()
	if err != nil {
		t.Errorf("Error getting Docker client: %v", err)
	}
	t.Log(client.ClientVersion())
}
