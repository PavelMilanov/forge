package template

import (
	"fmt"
	"os"

	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
	"github.com/spf13/cobra"
)

var (
	file  string
	vault *api.Vault
)

var TmpCmd = &cobra.Command{
	Use:       "tmpl [command]",
	Short:     "Manage template",
	Example:   "forge tmpl",
	ValidArgs: []string{"list"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	if err := os.MkdirAll(config.TEMPLATE_PATH, 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
