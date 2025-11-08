package deploy

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/spec"
	"github.com/spf13/cobra"
)

var (
	deploy bool
	config bool
)

var getCmd = &cobra.Command{
	Use:     "get [OPTIONS] [FLAGS]",
	Short:   "Get project information",
	Example: "forge deploy get <project> | forge deploy get <project> -p <param>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		secrets, err := vault.API.Get(ctx, args[0])
		if err != nil {
			errors.VaultErrors(err)
		}
		data := secrets.Data
		project, err := spec.NewSpec(data["mode"].(string))
		project.Parse(data)
		if err != nil {
			errors.VaultErrors(err)
		}
		if deploy {
			project.Print()
		} else if config {
			tmpl, exists := data["template"].(string)
			if !exists {
				errors.VaultErrors(fmt.Errorf("value not found"))
			}
			std, err := project.Generate(tmpl)
			if err != nil {
				errors.VaultErrors(err)
			}
			fmt.Print(std)
		}
	},
}

func init() {
	DeployCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&deploy, "deploy", "d", false, "project deploy secret")
	getCmd.Flags().BoolVarP(&config, "config", "c", false, "project config secret")
}
