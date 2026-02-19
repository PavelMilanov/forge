package agent

import (
	"github.com/PavelMilanov/forge/config"
)

type Agent interface {
	DeployStack(name string) error
	CreateStack(endpointId int, stackName string, template string) error
	ListEndpoints() (map[int]string, error)
}

func NewAgent() Agent {
	switch config.AppConfig.Agent.Type {
	case "portainer":
		return NewPortainerAgent()
	case "ssh":
		return NewSSHAgent()
	default:
		return nil
	}
}
