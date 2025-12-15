package deploy

import "github.com/spf13/cobra"

var DeployCmd = &cobra.Command{
	Use:   "deploy [command]",
	Short: "Manage deployment",
	// Example: "forge deploy",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
