package spec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"github.com/PavelMilanov/forge/config"
)

/*
Swarm интерфейс взаимодействия с docker swarm конфигурацией.
*/
type Swarm struct {
	Image    string `json:"image"`
	Tag      string `json:"tag"`
	Replicas int    `json:"replicas"`
}

/*
Init инициализирует модель с параметрами по умолчанию.
*/
func (s *Swarm) Init() {
	if s.Image == "" {
		s.Image = "image"
	}
	if s.Tag == "" {
		s.Tag = "latest"
	}
	if s.Replicas == 0 {
		s.Replicas = 1
	}
}

/*
Generate генерирует содержимое файла конфигурации на основе шаблона и данных модели.

Params:

	tmp - шаблон.

Returns:

	string - сгенерированное содержимое.
	err - ошибка, если она возникла.
*/
func (s *Swarm) Generate(tmp string) (string, error) {
	tmpl := template.Must(template.New("tmpl").Parse(tmp))
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, s)
	if err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}

/*
Parse преобразует входящие данные из Vault в структуру модели.

Params:

	data - входящие данные из Vault
*/
func (s *Swarm) Parse(data map[string]any) {
	deploy := data["deploy"].(map[string]any)
	switch v := deploy["replicas"].(type) {
	case json.Number:
		i, _ := strconv.Atoi(v.String())
		s.Replicas = i
	case string:
		i, _ := strconv.Atoi(v)
		s.Replicas = i
	}
	s.Image = deploy["image"].(string)
	s.Tag = deploy["tag"].(string)
}

/*
Update обновляет данные модели на основе входящих данных.

Params:

	data - входящие данные в формате key=value (флаги командной строки)

Returns:

	error - ошибка, если она возникла
*/
func (s *Swarm) Update(data []string) error {
	check := func(data []string) error {
		for _, param := range data {
			value := strings.Split(param, "=")
			if len(value) != 2 {
				return fmt.Errorf("Format is incorrect. Try: forge set <project> -p param=value")
			}
			found := false
			for _, flag := range config.SWARMPARAMS {
				if value[0] == flag {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("Unknown parameter: %s", value[0])
			}
		}
		return nil
	}
	if err := check(data); err != nil {
		return err
	}
	buf := make(map[string]string)
	for _, param := range data {
		value := strings.Split(param, "=")
		buf[value[0]] = value[1]
	}
	if len(buf["image"]) > 0 {
		s.Image = buf["image"]
	}
	if len(buf["tag"]) > 0 {
		s.Tag = buf["tag"]
	}
	if len(buf["replicas"]) > 0 {
		format, err := strconv.Atoi(buf["replicas"])
		if err != nil {
			return err
		}
		s.Replicas = format
	}
	return nil
}

/*
Print форматированный вывод данных модели в консоль.

Params:

	alias - алиас проекта
*/
func (s *Swarm) Print() {
	fmt.Printf(`image: %s
tag: %s
replicas: %d
`, s.Image, s.Tag, s.Replicas)
}
