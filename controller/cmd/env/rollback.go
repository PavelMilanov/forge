package env

import (
	"context"

	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/spec"
	"github.com/spf13/cobra"
)

var versionNum int

var rollbackCmd = &cobra.Command{
	Use:     "rollback [PROJECT] [FLAGS]",
	Short:   "Rollback config to version",
	Example: "forge env rollback [PROJECT] -v <version>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		secrets, err := vault.API.Rollback(ctx, args[0], versionNum)
		if err != nil {
			errors.VaultErrors(err)
		}
		data := secrets.Data
		project, err := spec.NewSpec(data["mode"].(string))
		project.Parse(data)
		if err != nil {
			errors.VaultErrors(err)
		}
		project.Print()
	},
}

func init() {
	EnvCmd.AddCommand(rollbackCmd)
	rollbackCmd.Flags().IntVarP(&versionNum, "version", "v", 0, "number version")
	rollbackCmd.MarkFlagRequired("version")
}
