package cmd

import (
	"os"

	"github.com/PavelMilanov/forge/cmd/deploy"
	"github.com/PavelMilanov/forge/cmd/env"
	"github.com/PavelMilanov/forge/cmd/template"
	"github.com/PavelMilanov/forge/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "forge",
	Short:   "cli-utility for managing ci/cd integration with infrastructure",
	Version: config.VERSION,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		env, err := config.NewEnv(config.FORGE_PATH, config.FORGE_FILE)
		if err != nil {
			return err
		}
		config.NewAppConfig(env)
		return nil
	},
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
	rootCmd.AddCommand(template.TmpCmd)
	rootCmd.AddCommand(deploy.DeployCmd)
}
