package docker

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/config"
	"github.com/compose-spec/compose-go/v2/loader"
	"github.com/compose-spec/compose-go/v2/types"
)

type Stack struct {
	App  *types.Project
	Mode int
}

func NewStack(file, projectName string) (*Stack, error) {
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
