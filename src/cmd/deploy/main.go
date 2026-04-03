package deploy

import (
	"github.com/PavelMilanov/forge/models"
	"github.com/spf13/cobra"
)

var (
	endpoints       []models.PortainerEndpoint // предзагрузка эндпоинтов из portainer
	endpointAliases []string                   // алиасы эндпоинтов для cli
)

var DeployCmd = &cobra.Command{
	Use:       "deploy [command]",
	Short:     "Развертывание и обновление стеков через агента",
	Long:      "Группа команд deploy работает с агентом развертывания (например, Portainer): позволяет просматривать доступные endpoint-ы и управлять стеками. Основной путь деплоя: deploy stack ... --file <path>.",
	Example:   "forge deploy stack prod-swarm -n admin -f ./docker-stack.yml --mode upsert",
	ValidArgs: []string{"stack", "list"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// init объявлен для единообразия структуры пакета deploy.
func init() {
}
