package cmd

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/spec"
	"github.com/spf13/cobra"
)

var params []string

var setCmd = &cobra.Command{
	Use:     "set [OPTIONS] [FLAGS]",
	Short:   "Set values of the project",
	Example: "forge set <project> -p tags=<string> -p replicas=<number>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(params) == 0 {
			errors.SpecErrors(fmt.Errorf("no parameters detected"))
		}
		ctx := context.Background()
		secrets, err := vault.API.Get(ctx, args[0])
		if err != nil {
			errors.VaultErrors(err)
		}
		value, exists := secrets.Data["deploy"]
		if !exists {
			errors.VaultErrors(fmt.Errorf("value not found"))
		}
		project, err := spec.NewSpec(secrets.Data["mode"].(string))
		if err != nil {
			errors.SpecErrors(err)
		}
		project.Parse(value.(map[string]any))
		if err := project.Update(params); err != nil {
			errors.SpecErrors(err)
		}
		_, err = vault.API.Patch(ctx, args[0], map[string]any{"deploy": project})
		if err != nil {
			errors.VaultErrors(err)
		}
		project.Print(args[0])

	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringSliceVarP(&params, "param", "p", []string{}, "project parameter")
}
