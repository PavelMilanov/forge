package env

import (
	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var (
	projectTemplate string
	projectMode     string
	vault           *api.Vault
)

var EnvCmd = &cobra.Command{
	Use:     "env",
	Short:   "Manage environment",
	Example: "forge env",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	var err error
	vault, err = api.NewVaultClient()
	if err != nil {
		errors.VaultErrors(err)
	}
	if err := vault.Login(); err != nil {
		errors.VaultErrors(err)
	}
}

func defaultFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&projectTemplate, "template", "t", "", "path to project/to/template.yml")
	cmd.Flags().StringVarP(&projectMode, "mode", "m", "compose", "project mode: compose | swarm | kubernetes")
	cmd.MarkFlagRequired("template")
	cmd.MarkFlagRequired("compose")
}
