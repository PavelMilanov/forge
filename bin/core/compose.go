package core

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/compose-spec/compose-go/v2/loader"
	"github.com/sirupsen/logrus"
)

type DockerCompose struct {
	Images     []string
	Containers []string
}

func (compose *DockerCompose) Parse(file string) error {
	ctx := context.Background()

	// Загружаем конфигурационные файлы.
	// Второй параметр – список файлов, а третий – рабочая директория.
	configDetails, err := loader.LoadConfigFiles(ctx, []string{file}, ".")
	if err != nil {
		return fmt.Errorf("error loading config files: %w", err)
	}

	// Парсим загруженную конфигурацию и получаем compose проект.
	project, err := loader.LoadWithContext(ctx, *configDetails, func(o *loader.Options) {
		if name, ok := o.GetProjectName(); !ok || name == "" {
			o.SetProjectName("default_project", true)
		}
	})
	if err != nil {
		return fmt.Errorf("error loading project: %w", err)
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
