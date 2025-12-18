package env

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/spec"
	"github.com/spf13/cobra"
)

var (
	config bool
)

/*
forge env get dev | grep 'tag:' | awk '{print $2}'
*/
var getCmd = &cobra.Command{
	Use:     "get [PROJECT] [FLAGS]",
	Short:   "Get project information",
	Example: "forge env get <project> | forge env get <project> -c",
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
		if config {
			tmpl, exists := data["template"].(string)
			if !exists {
				errors.VaultErrors(fmt.Errorf("value not found"))
			}
			std, err := project.Generate(tmpl)
			if err != nil {
				errors.VaultErrors(err)
			}
			fmt.Print(std)
		} else {
			project.Print()
		}
	},
}

func init() {
	EnvCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&config, "config", "c", false, "project config secret")
}
