package cmd

import (
	"os"

	"github.com/PavelMilanov/forge/cmd/deploy"
	"github.com/PavelMilanov/forge/cmd/env"
	"github.com/PavelMilanov/forge/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "forge",
	Short:   "cli-utility for managing ci/cd integration with infrastructure",
	Version: config.VERSION,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(env.EnvCmd)
	rootCmd.AddCommand(deploy.DeployCmd)
}
