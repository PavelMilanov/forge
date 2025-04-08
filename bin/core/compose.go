package core

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/compose-spec/compose-go/v2/loader"
	"github.com/compose-spec/compose-go/v2/types"
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
		return nil, fmt.Errorf("error getting absolute path: %w", err)
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
		return nil, fmt.Errorf("error loading project: %w", err)
	}
	return &Compose{
		App:  project,
		Dir:  filepath.Dir(path),
		File: path,
	}, nil
}

func (c *Compose) Deploy() error {
	//  вычислить образа
	// вычислить запущенные контейнеры
	// привести в соответсвие
	// проверить соостояние
	// 1. если нет запущенных контейнеров проекта - запустить
	// 2. если есть запущенные контейнеры проекта - проверить актуальность
	// 3. обновить
	// 4. проверить
	docker, err := NewDocker()
	if err != nil {
		return fmt.Errorf("error %w", err)
	}
	containers, err := docker.GetProjectContainers(c.Dir)
	if err != nil {
		return fmt.Errorf("error %w", err)
	}
	switch len(containers) {
	case 0:
		c.command("up")
		return nil
	default:
		fmt.Println("обновление не требуется")
		return nil
	}
}

func (c *Compose) command(command string) error {
	switch command {
	case "up":
		cmd := exec.Command("docker", "compose", "-f", c.File, "up", "-d")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error %w", err)
		}
		fmt.Println("Docker Compose успешно запущен")
		return nil
	case "down":
		cmd := exec.Command("docker", "compose", "-f", c.File, "down")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error %w", err)
		}
		fmt.Println("Docker Compose успешно остановлен")
		return nil
	default:
		return errors.New("неизвестная команда")
	}
}
