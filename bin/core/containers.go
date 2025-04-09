package core

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// Модель для взаимодействия с сущноснями Docker.
type Docker struct {
	Client *client.Client
}

// NewDocker инициализирует клиента docker.
func NewDocker() (*Docker, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &Docker{Client: cli}, nil
}

func (d *Docker) GetProjectContainers(project string) ([]container.Summary, error) {
	var projectInfo []container.Summary
	info, err := d.Client.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, container := range info {
		for _, label := range container.Labels {
			if label == project {
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

// PullImage скачивает образ из публичного/частного репозитория.
// Если не заданы логин и пароль(2 и 3 параметром) - скачивается публичный образ.
func (d *Docker) PullImage(name string, a ...string) error {
	var auth string
	options := image.PullOptions{}
	if len(a) != 0 {
		auth = base64.URLEncoding.EncodeToString(
			fmt.Appendf([]byte{}, "%s:%s", a[0], a[1]),
		)
		options = image.PullOptions{RegistryAuth: auth}
	}
	reader, err := d.Client.ImagePull(context.Background(), name, options)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	defer reader.Close()
	dec := json.NewDecoder(reader)
	for {
		var status map[string]interface{}
		if err := dec.Decode(&status); err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("error: %w", err)
		}
		if id, ok := status["id"]; ok {
			fmt.Printf("%s: ", id)
		}
		if s, ok := status["status"]; ok {
			fmt.Printf("%s", s)
		}
		if progress, ok := status["progress"]; ok {
			fmt.Printf(" %s", progress)
		}
	}
	fmt.Println(fmt.Sprintf("Образ %s загружен!", name))
	return nil
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
