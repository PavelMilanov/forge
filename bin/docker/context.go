package docker

import (
	"github.com/docker/docker/client"
	dockerHost "github.com/docker/go-sdk/context"
)

func GetDockerClient() (*client.Client, error) {
	// ctxName, err := dockerHost.Current()
	host, err := dockerHost.CurrentDockerHost()
	// host2, err := dockerCtx.DockerHostFromContext("my-context")
	// list, err := dockerHost.List()
	// info, err := dockerCtx.Inspect(ctxName)
	if err != nil {
		return nil, err
	}
	// fmt.Println(ctxName)
	// fmt.Println(list, host)

	cli, err := client.NewClientWithOpts(
		client.WithHost(host),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	// ctrs, err := cli.ContainerList(context.Background(), container.ListOptions{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Контейнеры:", ctrs)
	return cli, nil
}
