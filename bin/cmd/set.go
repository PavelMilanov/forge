/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set version for service",
	Long: `Set version for service. For example:

forge -f docker-compose.yml set version 1.0.0`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "version":
			fmt.Println("Setting version for service1 to", args[1])
		default:
			fmt.Println("Unknown parameter:", args[0])
		}
		fmt.Println("set called", args)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
