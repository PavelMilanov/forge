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
	Short: "Список эндпоинтов",
	Long:  "Позволяет получить список эндпоинтов.",
	Example: `Получить список доступных эндпоинтов:
forge deploy list
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := api.NewPortainer(
			config.AppConfig.Agent.Credentials.Url,
			config.AppConfig.Agent.Credentials.Key)
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

func init() {
	DeployCmd.AddCommand(listCmd)
}
