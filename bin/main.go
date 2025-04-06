package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	images, err := cli.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(images)
}
