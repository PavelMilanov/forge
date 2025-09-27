package cmd

import (
	"os"

	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var (
	projectTemplate string
	projectMode     string
	projectAlias    string
	vault           *api.VaultAPI
	hostPath        string
	hostAddr        string
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
	cmd.Flags().StringVarP(&projectTemplate, "file", "f", "", "path to project/to/template.yml")
	cmd.Flags().StringVarP(&projectMode, "mode", "m", "compose", "project mode: compose | swarm | kubernetes")
	cmd.Flags().StringVarP(&projectAlias, "alias", "a", "", "unique alias for the project")
	cmd.Flags().StringVarP(&hostPath, "path", "p", "/var/app", "path to remote host project directory")
	cmd.Flags().StringVarP(&hostAddr, "remote", "r", "", "remote host address")
	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("alias")
	cmd.MarkFlagRequired("remote")
}
