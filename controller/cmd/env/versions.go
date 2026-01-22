package env

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var versionsCmd = &cobra.Command{
	Use:     "versions [PROJECT]",
	Short:   "Print version history to project",
	Example: "forge env versions [PROJECT]",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		content, err := vault.API.GetVersionsAsList(ctx, args[0])
		if err != nil {
			errors.VaultErrors(err)
		}
		for _, metadata := range content {
			fmt.Printf(`Version: %d
Created: %s

`, metadata.Version, metadata.CreatedTime.Format("02/01/2006 15:04:05"))
		}
	},
}

func init() {
	EnvCmd.AddCommand(versionsCmd)
}
