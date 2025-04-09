package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/core"
)

func main() {
	switch os.Args[1] {
	case "deploy":
		deploy := flag.NewFlagSet("deploy", flag.ExitOnError)
		f := deploy.String("f", ".", "forge -f <path to docker-compose file>")
		deploy.Parse(os.Args[2:])
		compose, err := core.NewCompose(*f)
		if err != nil {
			fmt.Println(err)
		}
		compose.Deploy()
	case "version":
		cli, _ := core.NewDocker()
		info, _ := cli.Client.ServerVersion(context.Background())
		text := fmt.Sprintf("Forge : %s\nGo: %s\nlatfom: %s\nDocker: %s", config.VERSION, runtime.Version(), info.Platform.Name, info.Version)
		fmt.Println(text)
	default:
		fmt.Println("command not found")
	}

}
