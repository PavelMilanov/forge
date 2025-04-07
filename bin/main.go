package main

import (
	"flag"
	"fmt"
	"os"

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
	default:
		fmt.Println("command not found")
	}

}
