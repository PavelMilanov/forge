package deploy

import (
	"github.com/PavelMilanov/forge/models"
	"github.com/spf13/cobra"
)

var (
	endpoints       []models.PortainerEndpoint // предзагрузка эндпоинтов из portainer
	endpointAliases []string                   // алиасы эндпоинтов для cli
	deplyEndpoint   int                        // индекс выбранного эндпоинта для проброса в Portainer API
	deployStack     string
	deployTemplate  string
)

var DeployCmd = &cobra.Command{
	Use:       "deploy [command]",
	Short:     "Модуль развертывания",
	Long:      "Модуль развертывания позволяет создавать и развертывать стеки на основе шаблонов или обновлять существующие стеки через выбранного агента",
	Example:   "forge deploy",
	ValidArgs: []string{"stack", "list"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
}
