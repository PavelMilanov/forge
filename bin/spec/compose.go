package spec

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/PavelMilanov/forge/config"
)

/*
Compose - абстракция над docker compose спецификацией в файлах конфигурации.
*/
type Compose struct {
	Tag string `json:"tag"`
}

/*
Init инициализирует модель с параметрами по умолчанию.
*/
func (c *Compose) Init() {
	if c.Tag == "" {
		c.Tag = "latest"
	}
}

/*
Generate генерирует файл конфигурации на основе шаблона и данных модели.
*/
func (c *Compose) Generate(path, alias string) (string, error) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, c)
	if err != nil {
		return "", err
	}
	fileName := filepath.Join(config.CONFIG_PATH, fmt.Sprintf("%s-%s.yml", config.SPECMODE["compose"], alias))
	if err := os.WriteFile(fileName, buf.Bytes(), 0644); err != nil {
		return "", err
	}
	return fileName, nil
}

/*
Parse преобразует входящие данные из Vault в структуру модели.
*/
func (c *Compose) Parse(data map[string]any) {
	c.Tag = data["tag"].(string)
}

/*
Update обновляет данные модели на основе входящих данных.
*/
func (c *Compose) Update(data []string) error {
	if len(data) != 1 { // у спецификации compose 1 параметр
		return fmt.Errorf("1 parameter expected")
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
		c.Tag = buf["tag"]
	}
	return nil
}

/*
Print форматированный вывод данных модели в консоль.
*/
func (c *Compose) Print(alias string) {
	fmt.Printf(`%s
  tag: %s
`, alias, c.Tag)
}
