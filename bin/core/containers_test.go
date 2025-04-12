package core

import (
	"fmt"
	"testing"
	"time"
)

func TestGetProjectContainers(t *testing.T) {
	docker, _ := NewDocker()
	containers, err := docker.GetProjectContainers("test")
	if err != nil {
		t.Fatal(err)
	}
	for _, container := range containers {
		t.Logf("%+v", container)
	}
}

func TestPullImage(t *testing.T) {
	docker, err := NewDocker()
	if err != nil {
		t.Fatal(err)
	}
	images := []string{"nginx", "alpine"}
	start := time.Now()
	for _, image := range images {
		if err := docker.PullImage(image); err != nil {
			t.Fatal(err)
		}
	}
	fmt.Println(time.Since(start))
}
