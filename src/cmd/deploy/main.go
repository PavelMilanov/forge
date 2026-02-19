package deploy

import (
	"github.com/spf13/cobra"
)

var (
	endpoints       map[int]string // предзагрузка эндпоинтов из portainer
	endpointAliases []string       // алиасы эндпоинтов для cli
	deplyEndpoint   int            // индекс выбранного эндпоинта для проброса в Portainer API
	deployStack     string
	deployTemplate  string
)

var DeployCmd = &cobra.Command{
	Use:     "deploy [command]",
	Short:   "Модуль развертывания",
	Long:    "Модуль развертывания позволяет создавать и развертывать стеки на основе шаблонов или обновлять существующие стеки через выбранного агента",
	Example: "forge deploy",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
}
