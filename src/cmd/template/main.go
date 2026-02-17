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
	Use:       "tmpl [command]",
	Short:     "Модуль управления шаблонами проектов",
	Example:   "forge tmpl",
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

func init() {
	if err := os.MkdirAll(config.TEMPLATE_PATH, 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
