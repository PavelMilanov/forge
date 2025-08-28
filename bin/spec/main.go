package spec

import (
	"errors"

	"github.com/PavelMilanov/forge/config"
)

/*
NewSpec - спецификация данных инфраструктуры.
*/
type Spec interface {
	Init()
	Generate(path, alias string) (string, error)
	Parse(map[string]any)
	Update([]string) error
	Print(alias string)
}

func NewSpec(mode string) (Spec, error) {
	switch mode {
	case config.SPECMODE["swarm"]:
		spec := Swarm{}
		return &spec, nil
	case config.SPECMODE["compose"]:
		spec := Compose{}
		return &spec, nil
	case config.SPECMODE["kubernetes"]:
		spec := Kubernetes{}
		return &spec, nil
	default:
		return nil, errors.New("unknown mode")
	}
}
