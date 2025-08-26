package spec

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

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

func (c *Compose) New(path, alias string) error {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, c)
	if err != nil {
		return err
	}
	if err := os.WriteFile(config.SPECMODE["compose"]+"-"+alias+".yml", buf.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

/*
Parse преобразует входящие данные из Vault в форматированный вывод согласно модели.
*/
func (c *Compose) Parse(data map[string]any) {
	c.Tag = data["tag"].(string)
	fmt.Printf("%+v\n", *c)
}
