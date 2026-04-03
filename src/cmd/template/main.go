package template

import (
	"fmt"
	"os"

	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
	"github.com/spf13/cobra"
)

var appTemplate *api.Template

var TmpCmd = &cobra.Command{
	Use:       "templates [command]",
	Aliases:   []string{"tpl", "tmpl"},
	Short:     "Работа с шаблонами конфигурации проектов",
	Long:      "Группа команд templates предназначена для управления шаблонами, которые используются в env get -c для генерации итоговых YAML-файлов по данным из Vault.",
	Example:   "forge templates list",
	ValidArgs: []string{"list"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := api.NewTemplate()
		if err != nil {
			return err
		}
		appTemplate = cfg
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// init создает каталог шаблонов при старте пакета template.
func init() {
	if err := os.MkdirAll(config.TEMPLATE_PATH, 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
