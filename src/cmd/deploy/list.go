package deploy

import (
	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/text"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Показать доступные endpoint-ы агента",
	Long:  "Запрашивает список endpoint-ов у Portainer API и выводит их в табличном виде. Полученные имена endpoint-ов используются как аргумент команды deploy stack.",
	Example: `Получить список доступных эндпоинтов:
forge deploy list
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := api.NewPortainer(
			config.AppConfig.Agent.Credentials.Url,
			config.AppConfig.Agent.Credentials.Key,
			config.AppConfig.Agent.Credentials.Teams)
		if err != nil {
			errors.DeployErrors(err)
		}
		endpoints, err := cfg.GetEndpoints()
		if err != nil {
			errors.DeployErrors(err)
		}
		if err := text.PrintEndpoints(endpoints); err != nil {
			errors.DeployErrors(err)
		}
	},
}

// init регистрирует подкоманду deploy list.
func init() {
	DeployCmd.AddCommand(listCmd)
}
