package spec

import (
	"errors"
)

type Spec interface {
	Init() error
}

func NewSpec(mode string) (Spec, error) {
	switch mode {
	case "swarm":
		spec := Swarm{}
		return &spec, nil
	case "compose":
		spec := Compose{}
		return &spec, nil
	case "kubernetes":
		spec := Kubernetes{}
		return &spec, nil
	default:
		return nil, errors.New("unknown mode")
	}
}
