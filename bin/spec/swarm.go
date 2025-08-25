package spec

import (
	"bytes"
	"html/template"
	"os"

	"github.com/PavelMilanov/forge/config"
)

/*
Swarm - абстракция над docker swarm спецификацией в файлах конфигурации.
*/
type Swarm struct {
	Tag      string
	Replicas int
}

func (s *Swarm) Init(path, alias string) error {
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
