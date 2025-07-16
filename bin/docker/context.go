package docker

import (
	"errors"

	"github.com/docker/docker/client"
	dockerHost "github.com/docker/go-sdk/context"
)

func GetDockerClient(context ...string) (*client.Client, error) {
	switch len(context) {
	case 1:
		ctxName := context[0]
		host, err := dockerHost.DockerHostFromContext(ctxName)
		if err != nil {
			return nil, err
		}
		cli, err := client.NewClientWithOpts(
			client.WithHost(host),
			client.WithAPIVersionNegotiation(),
		)
		if err != nil {
			return nil, err
		}
		return cli, nil
	case 0:
		host, err := dockerHost.CurrentDockerHost()
		if err != nil {
			return nil, err
		}
		cli, err := client.NewClientWithOpts(
			client.WithHost(host),
			client.WithAPIVersionNegotiation(),
		)
		if err != nil {
			return nil, err
		}
		return cli, nil
	default:
		return nil, errors.New("docker context not found")
	}
}
