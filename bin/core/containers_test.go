package core

import (
	"testing"

	"github.com/docker/docker/client"
)

var cli, _ = client.NewClientWithOpts(client.FromEnv)

func TestGetProjectContainers(t *testing.T) {
	_, err := GetContainers(cli, ".")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetContainer(t *testing.T) {
	container, err := GetContainer(cli, "de300c34002ac8cf26238e6d13599bf04d230ca98cf515e2b07f53a2cd72d7b7")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", container)
}

func TestGetLogsContainer(t *testing.T) {
	logs, err := GetLogsContainer(cli, "de300c34002ac8cf26238e6d13599bf04d230ca98cf515e2b07f53a2cd72d7b7")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", string(logs))
}

func TestRestartContainer(t *testing.T) {
	err := RestartContainer(cli, "de300c34002ac8cf26238e6d13599bf04d230ca98cf515e2b07f53a2cd72d7b7")
	if err != nil {
		t.Fatal(err)
	}
}

func TestStopContainer(t *testing.T) {
	err := StopContainer(cli, "de300c34002ac8cf26238e6d13599bf04d230ca98cf515e2b07f53a2cd72d7b7")
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartContainer(t *testing.T) {
	err := StartContainer(cli, "de300c34002ac8cf26238e6d13599bf04d230ca98cf515e2b07f53a2cd72d7b7")
	if err != nil {
		t.Fatal(err)
	}
}
