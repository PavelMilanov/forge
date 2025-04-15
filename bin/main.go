package main

import (
	"fmt"
	"os"

	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/portainer"
)

func main() {
	env := config.NewEnv(config.CONFIG_PATH, "forge.yaml")
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "build":
			fmt.Println("build")
		case "deploy":
			// deploy := flag.NewFlagSet("deploy", flag.ExitOnError)
			// f := deploy.String("f", ".", "forge -f <path to docker-compose file>")
			// deploy.Parse(os.Args[2:])
			fmt.Println("deploy")
		case "update":
			fmt.Println("update command")
		case "environments":
			client := portainer.NewPortainerClient(env)
			if err := client.GetEnvironments(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case "version":
			fmt.Println(config.VERSION)
		default:
			fmt.Println("command not found")
		}
	}
}
