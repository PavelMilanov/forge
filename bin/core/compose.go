package core

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/PavelMilanov/forge/config"
	"github.com/compose-spec/compose-go/v2/loader"
	"github.com/compose-spec/compose-go/v2/types"
	"github.com/sirupsen/logrus"
)

// Модель для взаимодействия с сущноснями Docker Compose.
type Compose struct {
	App  *types.Project
	Dir  string
	File string
}

// NewCompose инициализирует docker-compose файл.
func NewCompose(file string) (*Compose, error) {
	path, err := filepath.Abs(file)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	project, err := loader.LoadWithContext(context.Background(), types.ConfigDetails{
		ConfigFiles: []types.ConfigFile{{
			Filename: path,
		}},
	},
		func(o *loader.Options) {
			if name, ok := o.GetProjectName(); !ok || name == "" {
				o.SetProjectName("forge", true)
			}
		})
	if err != nil {
		logrus.Errorln(err)
		return nil, fmt.Errorf("error loading project: %w", err)
	}
	return &Compose{
		App:  project,
		Dir:  filepath.Dir(path),
		File: path,
	}, nil
}

func (c *Compose) Deploy(config *config.Env) error {
	docker, err := NewDocker()
	if err != nil {
		return fmt.Errorf("error %w", err)
	}
	containers, err := docker.GetProjectContainers(c.Dir)
	if err != nil {
		return fmt.Errorf("error %w", err)
	}
	var composeImage []string
	for _, servise := range c.App.Services {
		composeImage = append(composeImage, servise.Image)
	}
	switch len(containers) {
	case 0:
		fmt.Println("Начинаю первичный запуск проекта")
		for _, item := range c.App.Services {
			if err := docker.PullImage(item.Name); err != nil {
				logrus.Error(err)
				return fmt.Errorf("error %w", err)
			}
		}
		if err := c.command("up"); err != nil {
			logrus.Error(err)
			return fmt.Errorf("error %w", err)
		}
		fmt.Println("Проект упешно запущен")
		return nil
	default:
		var currentImage []string
		for _, container := range containers {
			currentImage = append(currentImage, container.Image)
		}
		if len(composeImage) > len(currentImage) {
			fmt.Println("Обнаружены новые сервисы")
			var updateImage []string
			for _, image := range composeImage {
				if !slices.Contains(currentImage, image) {
					updateImage = append(updateImage, image)
				}
			}
			for _, item := range updateImage {
				if err := docker.PullImage(item); err != nil {
					return fmt.Errorf("error %w", err)
				}
			}
			c.command("up")
			fmt.Println("Проект упешно обновлен")
		} else if len(composeImage) < len(currentImage) {
			fmt.Println("обнаружено удаление сервисов")
			var updateImage []string
			for _, image := range currentImage {
				if !slices.Contains(composeImage, image) {
					updateImage = append(updateImage, image)
				}
			}
			containers, err := docker.GetProjectContainers(c.App.WorkingDir)
			if err != nil {
				return err
			}
			for _, container := range containers {
				if slices.Contains(updateImage, container.Image) {
					fmt.Println("Необходимо удалить контейнер", container.Names[0])
					if err := docker.DeleteContainer(container.ID); err != nil {
						return err
					}
					fmt.Println("Контейнер удален")
				}
			}
		} else {
			for _, item := range c.App.Services {
				if !slices.Contains(currentImage, item.Image) {
					if err := docker.PullImage(item.Image); err != nil {
						return fmt.Errorf("error %w", err)
					}
					c.command("down", item.Name)
					c.command("up", item.Name)
					fmt.Println("Проект упешно обновлен")
				}
			}
			fmt.Println("Обновление не требуется")
		}
		return nil
	}
}

func (c *Compose) command(command string, a ...string) error {
	var cmd *exec.Cmd
	switch command {
	case "up":
		if len(a) == 0 {
			cmd = exec.Command("docker", "compose", "-f", c.File, "up", "-d")
		} else {
			cmd = exec.Command("docker", "compose", "-f", c.File, "up", "-d", strings.Join(a, " "))
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error %w", err)
		}
		return nil
	case "down":
		if len(a) == 0 {
			cmd = exec.Command("docker", "compose", "-f", c.File, "down")
		} else {
			cmd = exec.Command("docker", "compose", "-f", c.File, "down", strings.Join(a, " "))
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error %w", err)
		}
		return nil
	default:
		return errors.New("неизвестная команда")
	}
}
