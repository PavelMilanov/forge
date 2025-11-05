package deploy

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var (
	deploy bool
	mode   bool
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
		if deploy {
			value, exists := secrets.Data["deploy"]
			if !exists {
				errors.VaultErrors(fmt.Errorf("value not found"))
			}
			fmt.Println(value.(map[string]any))
		} else if mode {
			value, exists := secrets.Data["mode"].(string)
			if !exists {
				errors.VaultErrors(fmt.Errorf("value not found"))
			}
			fmt.Println(value)
		} else if config {
			value, exists := secrets.Data["template"].(string)
			if !exists {
				errors.VaultErrors(fmt.Errorf("value not found"))
			}
			fmt.Println(value)
		}
	},
}

func init() {
	DeployCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&deploy, "deploy", "d", false, "project deploy secret")
	getCmd.Flags().BoolVarP(&mode, "mode", "m", false, "project mode secret")
	getCmd.Flags().BoolVarP(&config, "config", "c", false, "project config secret")
}
