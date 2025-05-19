package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/docker"
	"github.com/PavelMilanov/forge/utils"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Обновление сервиса до указанной версии",
	Long: `Обновление сервиса до указанной версии и генерация конфигурационного файла docker. Пример:

forge -f docker-compose.yml -s alpine update 1.0.0`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		stack, _ := docker.NewStack(dockerFile)
		data := map[string]interface{}{}
		data[stack.App.Name+"_"+dockerService] = args[0]
		_, err := vault.Patch(ctx, config.VAULT_PATH, data)
		if err != nil {
			os.Exit(1)
		}
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
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().StringVarP(&dockerService, "service", "s", "", "forge -s <service_name>")
	updateCmd.MarkPersistentFlagRequired("service")
}
