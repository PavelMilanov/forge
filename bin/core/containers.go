package core

import (
	"bytes"
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func GetContainers(client *client.Client, file string) ([]container.Summary, error) {
	var projectInfo []container.Summary
	info, err := client.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, container := range info {
		for _, label := range container.Labels {
			if label == file {
				projectInfo = append(projectInfo, container)
			}
		}
	}

	return projectInfo, nil
}

func GetContainer(client *client.Client, id string) (*container.ContainerJSONBase, error) {
	container, err := client.ContainerInspect(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return container.ContainerJSONBase, nil
}

func StartContainer(client *client.Client, id string) error {
	return client.ContainerStart(context.Background(), id, container.StartOptions{})
}

func StopContainer(client *client.Client, id string) error {
	return client.ContainerStop(context.Background(), id, container.StopOptions{})
}

func RestartContainer(client *client.Client, id string) error {
	return client.ContainerRestart(context.Background(), id, container.StopOptions{})
}

// GetLogsContainer
func GetLogsContainer(client *client.Client, id string) ([]byte, error) {
	data, err := client.ContainerLogs(context.Background(), id, container.LogsOptions{ShowStdout: true, ShowStderr: true, Follow: false,
		Timestamps: false})
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var stdoutBuf, stderrBuf bytes.Buffer
	_, err = stdcopy.StdCopy(&stdoutBuf, &stderrBuf, data)
	if err != nil {
		return nil, err
	}
	combined := append(stdoutBuf.Bytes(), stderrBuf.Bytes()...)
	return combined, nil
}
