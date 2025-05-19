package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/utils"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		secrets, err := vault.Get(ctx, config.VAULT_PATH)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tags := map[string]string{}
		for key, value := range secrets.Data {
			tags[key] = value.(string)
		}
		if err := utils.GenerateAppConfig(dockerFile, tags); err != nil {
			fmt.Println("Ошибка генерации конфигурационного файла:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
