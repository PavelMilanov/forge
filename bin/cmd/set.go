/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set version of the specified service",
	Example: `
forge set -a stage -s alpine 3.21
<stage-stack.yml>
services:
  alpine:
    image: alpine:3.21
    container_name: alpine
    restart: unless-stopped
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// ctx := context.Background()
		// data := map[string]interface{}{}
		// data[dockerService] = args[0]
		// _, err := vault.KV.Patch(ctx, dockerAlias, data)
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
		// text := fmt.Sprintf("%s version updated: %s", dockerService, args[0])
		// fmt.Println(text)
		// os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	addAliasFlags(setCmd)
}
