package env

import (
	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var (
	projectTemplate string
	projectMode     string
	projectAlias    string
	vault           *api.VaultAPI
	template        string
)

var EnvCmd = &cobra.Command{
	Use:   "env [command]",
	Short: "Manage environment",
	// Example: "forge env",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
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
	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("alias")

}
