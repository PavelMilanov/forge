package spec

import (
	"bytes"
	"html/template"
	"os"

	"github.com/PavelMilanov/forge/config"
)

/*
Compose - абстракция над docker compose спецификацией в файлах конфигурации.
*/
type Compose struct {
	Tag string
}

func (c *Compose) Init(path, alias string) error {
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
