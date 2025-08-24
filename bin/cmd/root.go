package cmd

import (
	"os"

	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/utils"
	"github.com/spf13/cobra"
)

var (
	// project       *docker.Stack
	dockerFile    string
	dockerService string
	dockerAlias   string
	vault         *utils.VaultClient
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
	vault = utils.NewVaultClient()
}

func addDefaultFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&dockerFile, "file", "f", "", "path to dockerProject.yml")
	cmd.Flags().StringVarP(&dockerAlias, "alias", "a", "", "unique project name")
	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("alias")
}

func addAliasFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&dockerAlias, "alias", "a", "", "project name")
	cmd.Flags().StringVarP(&dockerService, "service", "s", "", "service of project")
	cmd.MarkFlagRequired("alias")
	cmd.MarkFlagRequired("service")
}
