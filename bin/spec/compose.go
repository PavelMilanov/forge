package spec

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
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
Parse преобразует входящие данные из Vault в структуру модели.
*/
func (c *Compose) Parse(data map[string]any) {
	c.Tag = data["tag"].(string)
}
func (c *Compose) Update(data []string) {
	if len(data) != 1 { // у спецификации compose 1 параметр
		fmt.Println("Ошибка: ожидается один параметр")
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
	c.Tag = buf["tag"]
}

func (c *Compose) Print(alias string) {
	fmt.Printf(`%s
  tag: %s
`, alias, c.Tag)
}
