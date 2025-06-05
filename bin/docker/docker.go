// Package docker реализует функции для взаимодействия с Docker.
package docker

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/PavelMilanov/forge/config"
	"github.com/compose-spec/compose-go/v2/loader"
	"github.com/compose-spec/compose-go/v2/types"
)

type Stack struct {
	App  *types.Project
	Mode int
}

func NewStack(file, projectName string) (*Stack, error) {
	// dirs := strings.Split(filepath.Dir(file), "/")
	// fmt.Println(dirs)
	// projectName := dirs[len(dirs)-1]

	project, err := loader.LoadWithContext(context.Background(), types.ConfigDetails{
		ConfigFiles: []types.ConfigFile{{
			Filename: file,
		}},
	},
		func(o *loader.Options) {
			if name, ok := o.GetProjectName(); !ok || name == "" {
				o.SetProjectName(projectName, true)
			}
		})
	if err != nil {
		return nil, fmt.Errorf("error loading project: %w", err)
	}
	var stack Stack
	stack.App = project
	for _, service := range project.Services {
		if service.Deploy == nil {
			stack.Mode = config.DOCKERMOD["compose"]
		} else {
			stack.Mode = config.DOCKERMOD["stack"]
		}
	}
	return &stack, nil
}

func DockerCommand(command, env, filepath string, a ...string) error {
	var cmd *exec.Cmd
	switch command {
	case "up":
		if len(a) == 0 {
			cmd = exec.Command("docker", "--context", env, "compose", "-f", filepath, "up", "-d")
		} else {
			cmd = exec.Command("docker", "--context", env, "compose", "-f", filepath, "up", "-d", strings.Join(a, " "))
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error %w", err)
		}
		return nil
	case "update":
		if len(a) == 0 {
			cmd = exec.Command("docker", "--context", env, "compose", "-f", filepath, "up", "-d", "--force-recreate")
		} else {
			cmd = exec.Command("docker", "--context", env, "compose", "-f", filepath, "up", "-d", strings.Join(a, " "), "--force-recreate")
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error %w", err)
		}
		return nil
	case "down":
		if len(a) == 0 {
			cmd = exec.Command("docker", "--context", env, "compose", "-f", filepath, "down")
		} else {
			cmd = exec.Command("docker", "--context", env, "compose", "-f", filepath, "down", strings.Join(a, " "))
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error %w", err)
		}
		return nil
	default:
		return errors.New("unknown command")
	}
}

func RegistryLogin(env string, login, password, registry string) error {
	cmd := exec.Command("echo", password, "|", "docker", "--context", env, "login", "-u", login, "--password-stdin", registry)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error %w", err)
	}
	return nil
}
