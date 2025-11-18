package spec

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/PavelMilanov/forge/config"
)

/*
Compose	интерфейс взаимодействия с docker compose конфигурацией.
*/
type Compose struct {
	Image string `json:"image"`
	Tag   string `json:"tag"`
}

/*
Init инициализирует модель с параметрами по умолчанию.
*/
func (c *Compose) Init() {
	if c.Image == "" {
		c.Image = "image"
	}
	if c.Tag == "" {
		c.Tag = "latest"
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
func (c *Compose) Generate(tmp string) (string, error) {
	tmpl, err := template.ParseFiles(filepath.Join(config.TEMPLATE_PATH, tmp))
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, c); err != nil {
		return "", err
	}
	return buf.String(), nil
}

/*
Parse преобразует входящие данные из Vault в структуру модели.

Params:

	data - входящие данные из Vault
*/
func (c *Compose) Parse(data map[string]any) {
	deploy := data["deploy"].(map[string]any)
	c.Image = deploy["image"].(string)
	c.Tag = deploy["tag"].(string)
}

/*
Update обновляет данные модели на основе входящих данных.

Params:

	data - входящие данные в формате key=value (флаги командной строки)

Returns:

	error - ошибка, если она возникла
*/
func (c *Compose) Update(data []string) error {
	check := func(data []string) error {
		for _, param := range data {
			value := strings.Split(param, "=")
			if len(value) != 2 {
				return fmt.Errorf("Format is incorrect. Try: forge set <project> -p param=value")
			}
			found := false
			for _, flag := range config.COMPOSEPARAMS {
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
	if len(buf["tag"]) > 0 {
		c.Tag = buf["tag"]
	}
	if len(buf["image"]) > 0 {
		c.Image = buf["image"]
	}
	return nil
}

/*
Print форматированный вывод данных модели в консоль.
*/
func (c *Compose) Print() {
	fmt.Printf(`image: %s
tag: %s
`, c.Image, c.Tag)
}
