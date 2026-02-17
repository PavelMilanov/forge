package template

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Вывод шаблонов для проекта",
	Example: "forge tmpl list",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appTemplate.GetTemplates()
	},
}

func init() {
	TmpCmd.AddCommand(listCmd)
}
