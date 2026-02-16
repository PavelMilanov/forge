package deploy

import (
	"github.com/spf13/cobra"
)

var (
	deployStack    string
	deployTemplate string
)

var DeployCmd = &cobra.Command{
	Use:       "deploy [command]",
	Short:     "Модуль развертывания",
	Long:      "Модуль развертывания позволяет создавать и развертывать стеки на основе шаблонов или обновлять существующие стеки через выбранного агента",
	Example:   "forge deploy",
	ValidArgs: []string{"stack"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
}
