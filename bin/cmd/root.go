package cmd

import (
	"os"

	"github.com/PavelMilanov/forge/docker"
	"github.com/PavelMilanov/forge/utils"
	"github.com/spf13/cobra"
)

var (
	project       *docker.Stack
	dockerFile    string
	dockerService string
	dockerEnv     string
	vault         *utils.VaultClient
)

var rootCmd = &cobra.Command{
	Use:   "forge",
	Short: "cli-utility for managing ci/cd integration with docker infrastructure",
	Long: `cli utility for managing ci/cd integration with docker infrastructure.
`,
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
