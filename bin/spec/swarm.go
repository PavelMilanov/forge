package spec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PavelMilanov/forge/config"
)

/*
Swarm интерфейс взаимодействия с docker swarm конфигурацией.
*/
type Swarm struct {
	Image    string `json:"image,omitempty"`
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
Generate генерирует файл конфигурации на основе шаблона и данных модели.

Params:

	path - путь к шаблону
	alias - алиас для идентификации файла

Returns:

	fileName - путь к сгенерированному файлу
	err - ошибка, если она возникла
*/
func (s *Swarm) Generate(path, alias string) (string, error) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, s)
	if err != nil {
		return "", err
	}
	fileName := filepath.Join(config.CONFIG_PATH, fmt.Sprintf("%s-%s.yml", config.SPECMODE["swarm"], alias))
	if err := os.WriteFile(fileName, buf.Bytes(), 0644); err != nil {
		return "", err
	}
	return fileName, nil
}

/*
Parse преобразует входящие данные из Vault в структуру модели.

Params:

	data - входящие данные из Vault
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
	s.Image = data["image"].(string)
	s.Tag = data["tag"].(string)
}

/*
Update обновляет данные модели на основе входящих данных.

Params:

	data - входящие данные в формате key=value (флаги командной строки)

Returns:

	error - ошибка, если она возникла
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
func (s *Swarm) Print(alias string) {
	fmt.Printf(`%s
  image: %s
  tag: %s
  replicas: %d
`, alias, s.Image, s.Tag, s.Replicas)
}
