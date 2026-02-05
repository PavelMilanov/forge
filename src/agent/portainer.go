package agent

import (
	"fmt"

	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
)

type PortainerAgent struct {
	portainer *api.Portainer
}

func NewPortainerAgent() *PortainerAgent {
	return &PortainerAgent{}
}

func (a *PortainerAgent) DeployStack(name string) error {
	portainer, err := api.NewPortainer(
		config.AppConfig.Agent.Credentials.Url,
		config.AppConfig.Agent.Credentials.Key)
	if err != nil {
		return err
	}
	stacks, err := portainer.GetStacks()
	if err != nil {
		return err
	}
	for _, stack := range stacks {
		if stack.StackName == name {
			project, err := portainer.GetStackFile(stack)
			if err != nil {
				return err
			}
			resp, err := portainer.UpdateStack(project)
			if err != nil {
				return err
			}
			fmt.Println(resp)
			return nil
		}
	}
	return nil
}
