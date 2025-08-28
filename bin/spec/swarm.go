package spec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"strconv"
	"strings"

	"github.com/PavelMilanov/forge/config"
)

/*
Swarm - абстракция над docker swarm спецификацией в файлах конфигурации.
*/
type Swarm struct {
	Tag      string `json:"tag"`
	Replicas int    `json:"replicas"`
}

/*
Init инициализирует модель с параметрами по умолчанию.
*/
func (s *Swarm) Init() {
	if s.Tag == "" {
		s.Tag = "latest"
	}
	if s.Replicas == 0 {
		s.Replicas = 1
	}
}

func (s *Swarm) New(path, alias string) error {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, s)
	if err != nil {
		return err
	}
	if err := os.WriteFile(config.SPECMODE["swarm"]+"-"+alias+".yml", buf.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

/*
Parse преобразует входящие данные из Vault в структуру модели.
*/
func (s *Swarm) Parse(data map[string]any) {
	switch v := data["replicas"].(type) {
	case json.Number:
		i, _ := strconv.Atoi(v.String())
		s.Replicas = i
	case string:
		i, _ := strconv.Atoi(v)
		s.Replicas = i
	}
	s.Tag = data["tag"].(string)
}

/*
Update обновляет данные модели на основе входящих данных.
*/
func (s *Swarm) Update(data []string) error {
	if len(data) > 2 { // у спецификации swarm [1 или 2] параметра
		return fmt.Errorf("1 or 2 parameters expected")
	}
	buf := make(map[string]string)
	for _, param := range data {
		value := strings.Split(param, "=")
		if len(value) != 2 {
			return fmt.Errorf("Format is incorrect. Try: forge set <project> -p param=value")
		}
		buf[value[0]] = value[1]
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
*/
func (s *Swarm) Print(alias string) {
	fmt.Printf(`%s
  tag: %s
  replicas: %d
`, alias, s.Tag, s.Replicas)
}
