package spec

import (
	"errors"

	"github.com/PavelMilanov/forge/config"
)

/*
NewSpec спецификация данных инфраструктуры.
*/
type Spec interface {
	Init()
	Generate(path, alias string) (string, error)
	Parse(map[string]any)
	Update([]string) error
	Print(alias string)
}

/*
NewSpec инициализирует спецификацию данных инфраструктуры в зависимости от модели конфигурации.

Params:

	mode - модель конфигурации

Returns:

	Spec - спецификация данных инфраструктуры
	error - ошибка инициализации спецификации
*/
func NewSpec(mode string) (Spec, error) {
	switch mode {
	case config.SPECMODE["swarm"]:
		var spec Swarm
		return &spec, nil
	case config.SPECMODE["compose"]:
		var spec Compose
		return &spec, nil
	case config.SPECMODE["kubernetes"]:
		var spec Kubernetes
		return &spec, nil
	default:
		return nil, errors.New("unknown mode")
	}
}
