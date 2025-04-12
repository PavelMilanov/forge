package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/core"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetLevel(logrus.TraceLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006/01/02 15:04:00",
	})
}

func main() {
	log.Out = os.Stdout
	file, err := os.OpenFile("forge.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	}
	defer file.Close()
	switch os.Args[1] {
	case "deploy":
		deploy := flag.NewFlagSet("deploy", flag.ExitOnError)
		f := deploy.String("f", ".", "forge -f <path to docker-compose file>")
		deploy.Parse(os.Args[2:])
		compose, err := core.NewCompose(*f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := compose.Deploy(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	case "version":
		cli, _ := core.NewDocker()
		info, _ := cli.Client.ServerVersion(context.Background())
		text := fmt.Sprintf("Forge : %s\nGo: %s\nPlatform: %s\nDocker: %s", config.VERSION, runtime.Version(), info.Platform.Name, info.Version)
		fmt.Println(text)
	default:
		fmt.Println("command not found")
	}

}
