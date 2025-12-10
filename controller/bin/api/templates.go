package api

import (
	"fmt"
	"os"

	"github.com/PavelMilanov/forge/config"
)

func GetTemplates() {
	entries, err := os.ReadDir(config.TEMPLATE_PATH)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range entries {
		// Выводим имя файла или директории
		fmt.Println(entry.Name())
	}
}
