/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/spec"
	"github.com/spf13/cobra"
)

var param string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [OPTIONS]",
	Short: "Get project information",
	Example: `
forge get <project>

forge get <project> -p <param>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		secrets, err := vault.API.Get(ctx, args[0])
		if err != nil {
			errors.VaultErrors(err)
		}
		value, exists := secrets.Data["deploy"]
		if !exists {
			errors.VaultErrors(fmt.Errorf("value not found"))
		}
		if param != "" {
			fmt.Println(value.(map[string]any)[param])
		} else {
			project, err := spec.NewSpec(secrets.Data["mode"].(string))
			if err != nil {
				errors.SpecErrors(err)
			}
			project.Parse(secrets.Data["deploy"].(map[string]any))
			project.Print(args[0])
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	addAliasFlags(getCmd)
	getCmd.Flags().StringVarP(&param, "param", "p", "", "project parameter")
}
