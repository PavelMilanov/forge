package cmd

import (
	"os"

	"github.com/PavelMilanov/forge/docker"
	"github.com/PavelMilanov/forge/utils"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
)

var (
	project       *docker.Stack
	dockerFile    string
	dockerService string
	dockerEnv     string
	vault         *api.KVv2
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
	rootCmd.PersistentFlags().StringVarP(&dockerFile, "file", "f", "", "forge -f <docker-compose.yml|docker-stack.yml>")
	rootCmd.PersistentFlags().StringVarP(&dockerEnv, "env", "e", "default", "forge -e <env_name>")
	rootCmd.MarkPersistentFlagRequired("file")

	client := utils.NewVault()
	vault = client
}
