/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get version of the specified service",
	Example: `
forge get -a stage -s alpine
3.21
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		secrets, err := vault.KV.Get(ctx, dockerAlias)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if dockerService != "" {
			value, exists := secrets.Data[dockerService]
			if !exists {
				fmt.Println("Service not found")
				os.Exit(1)
			}
			fmt.Println(value.(string))
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	addAliasFlags(getCmd)
}
