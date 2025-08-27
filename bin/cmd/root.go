package cmd

import (
	"os"

	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var (
	projectFile  string
	projectMode  string
	projectAlias string
	vault        *api.VaultAPI
)

var rootCmd = &cobra.Command{
	Use:     "forge",
	Short:   "cli-utility for managing ci/cd integration with docker infrastructure",
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
	var err error
	vault, err = api.NewVaultClient()
	if err != nil {
		errors.VaultErrors(err)
	}
	vault.Set()
	if err := vault.RenewToken(); err != nil {
		errors.VaultErrors(err)
	}
}

func defaultFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&projectFile, "file", "f", "", "path to project/to/configuration.yml")
	cmd.Flags().StringVarP(&projectMode, "mode", "m", "compose", "project mode")
	cmd.Flags().StringVarP(&projectAlias, "alias", "a", "", "unique project name")
	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("alias")
}

func addAliasFlags(cmd *cobra.Command) {
	// cmd.Flags().StringVarP(&proAlias, "alias", "a", "", "project name")
	// cmd.Flags().StringVarP(&dockerService, "service", "s", "", "service of project")
	// cmd.MarkFlagRequired("alias")
	// cmd.MarkFlagRequired("service")
}
