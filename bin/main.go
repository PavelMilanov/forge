package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/core"
	"github.com/sirupsen/logrus"
)

func main() {
	logFile, err := os.OpenFile(filepath.Join(config.LOG_PATH, "forge.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(mw)
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006/01/02 15:04:00",
	})

	env := config.NewEnv(config.CONFIG_PATH, "forge.yaml")

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
		if err := compose.Deploy(env); err != nil {
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
