package template

import (
	"github.com/PavelMilanov/forge/api"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Set values of the project",
	Example: "forge tmpls list",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		api.GetTemplates()
	},
}

func init() {
	TmpCmd.AddCommand(listCmd)
}
