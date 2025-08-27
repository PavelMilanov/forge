package spec

import (
	"errors"

	"github.com/PavelMilanov/forge/config"
)

type Spec interface {
	Init()
	New(path, alias string) error
	Parse(map[string]any)
	Update([]string)
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
