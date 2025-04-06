package core

import (
	"context"
	"errors"
	"os"
	"os/exec"

	"github.com/compose-spec/compose-go/v2/loader"
	"github.com/compose-spec/compose-go/v2/types"
	"github.com/sirupsen/logrus"
)

type DockerCompose struct {
	Images     []string
	Containers []string
}

func (compose *DockerCompose) Parse(file string) error {
	project, err := loader.LoadWithContext(
		context.Background(),
		types.ConfigDetails{
			ConfigFiles: []types.ConfigFile{{
				Filename: file,
			}},
			Environment: nil,
		},
	)
	if err != nil {
		logrus.Error(err)
		return err
	}

	for _, service := range project.Services {
		compose.Images = append(compose.Images, service.Image)
		if service.ContainerName == "" {
			logrus.Warnf("у сервиса %s не указано имя контейнера", service.Name)
			continue
		}
		compose.Containers = append(compose.Containers, service.ContainerName)
	}
	return nil
}

func (compose *DockerCompose) Command(file, command string) error {
	switch command {
	case "up":
		cmd := exec.Command("docker", "compose", "-f", file, "up", "-d")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			logrus.Errorf("Ошибка запуска Compose: %v", err)
			return err
		}
		logrus.Info("Docker Compose успешно запущен")
		return nil
	case "down":
		cmd := exec.Command("docker", "compose", "-f", file, "down")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			logrus.Errorf("Ошибка остановки Compose: %v", err)
			return err
		}
		logrus.Info("Docker Compose успешно остановлен")
		return nil
	default:
		logrus.Errorf("Неизвестная команда: %s", command)
		return errors.New("неизвестная команда")
	}

}
