package core

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func GetContainers(client *client.Client) ([]container.Summary, error) {
	info, err := client.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return nil, err
	}

	return info, nil
}
