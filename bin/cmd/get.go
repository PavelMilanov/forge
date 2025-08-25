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

var param string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get project information",
	Example: `
forge get <project>
Output:
<key>:<value>

forge get <project> -p <param>
Output:
<value>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		secrets, err := vault.KV.Get(ctx, args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// project, err := spec.NewSpec(secrets.Data["mode"].(string))
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
		value, exists := secrets.Data["deploy"]
		if !exists {
			fmt.Println("value not found")
			os.Exit(1)
		}
		if param != "" {
			fmt.Println(value.(map[string]interface{})[param])
		} else {
			for key, val := range value.(map[string]interface{}) {
				text := fmt.Sprintf("%s:%s", key, val)
				fmt.Println(text)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	addAliasFlags(getCmd)
	getCmd.Flags().StringVarP(&param, "param", "p", "", "get project parameter")
}
