// Package docker реализует функции для взаимодействия с Docker.
package docker

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/PavelMilanov/forge/config"
	"github.com/compose-spec/compose-go/v2/loader"
	"github.com/compose-spec/compose-go/v2/types"
)

type Stack struct {
	App  *types.Project
	Dir  string
	Mode int
}

func NewStack(file string) (*Stack, error) {
	dirs := strings.Split(filepath.Dir(file), "/")
	projectName := dirs[len(dirs)-1]

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
	stack.Dir = filepath.Dir(file)
	for _, service := range project.Services {
		if service.Deploy == nil {
			stack.Mode = config.DOCKERMOD["compose"]
		} else {
			stack.Mode = config.DOCKERMOD["stack"]
		}
	}
	return &stack, nil
}

// func (d *Docker) command(command string, a ...string) error {
// 	var cmd *exec.Cmd
// 	switch command {
// 	case "up":
// 		if len(a) == 0 {
// 			cmd = exec.Command("docker", "compose", "-f", c.File, "up", "-d")
// 		} else {
// 			cmd = exec.Command("docker", "compose", "-f", c.File, "up", "-d", strings.Join(a, " "))
// 		}
// 		cmd.Stdout = os.Stdout
// 		cmd.Stderr = os.Stderr
// 		if err := cmd.Run(); err != nil {
// 			return fmt.Errorf("error %w", err)
// 		}
// 		return nil
// 	case "down":
// 		if len(a) == 0 {
// 			cmd = exec.Command("docker", "compose", "-f", c.File, "down")
// 		} else {
// 			cmd = exec.Command("docker", "compose", "-f", c.File, "down", strings.Join(a, " "))
// 		}
// 		cmd.Stdout = os.Stdout
// 		cmd.Stderr = os.Stderr
// 		if err := cmd.Run(); err != nil {
// 			return fmt.Errorf("error %w", err)
// 		}
// 		return nil
// 	default:
// 		return errors.New("неизвестная команда")
// 	}
// }
