package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/docker"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		stack, _ := docker.NewStack(dockerFile)
		data := map[string]interface{}{}
		for _, service := range stack.App.Services {
			data[stack.App.Name+"_"+service.Name] = "latest"
		}
		_, err := vault.Put(ctx, config.VAULT_PATH, data)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println("Проект инициализирован успешно")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
