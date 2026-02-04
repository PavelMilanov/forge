package api

import (
	"fmt"
	"os"

	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/errors"
)

type Template struct {
}

func NewTemplate() (*Template, error) {
	return &Template{}, nil
}

func (t *Template) GetTemplates() {
	entries, err := os.ReadDir(config.TEMPLATE_PATH)
	if err != nil {
		errors.ForgeErrors(err)
	}
	for _, entry := range entries {
		// Выводим имя файла или директории
		fmt.Println(entry.Name())
	}
}
