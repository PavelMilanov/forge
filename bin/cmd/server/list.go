/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, server := range env.SSH.Servers {
			fmt.Printf(
				`- %s
  user: %s
  server: %s
`, server.Host, server.User, server.Server)
		}
	},
}

func init() {
	ServerCmd.AddCommand(listCmd)
}
