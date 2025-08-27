package spec

import (
	"bytes"
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
	s.Tag = data["tag"].(string)
	s.Replicas = data["replicas"].(int)
}

func (s *Swarm) Update(data []string) {
	if len(data) != 2 { // у спецификации swarm 2 параметра
		fmt.Println("Ошибка: ожидается два параметра")
		os.Exit(1)
	}
	buf := make(map[string]string)
	for _, param := range data {
		value := strings.Split(param, "=")
		if len(value) != 2 {
			fmt.Println("Format is incorrect. Try: forge set <project> -p param=value")
			os.Exit(1)
		}
		buf[value[0]] = value[1]
	}
	s.Tag = buf["tag"]
	format, err := strconv.Atoi(buf["replicas"])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s.Replicas = format
}

func (s *Swarm) Print(alias string) {
	fmt.Printf(`%s
  tag: %s
  replicas: %d
`, alias, s.Tag, s.Replicas)
}
