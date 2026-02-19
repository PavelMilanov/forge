package deploy

import (
	"fmt"

	"github.com/PavelMilanov/forge/agent"
	"github.com/PavelMilanov/forge/errors"
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
		cfg := agent.NewAgent()
		endpoints, err := cfg.ListEndpoints()
		if err != nil {
			errors.DeployErrors(err)
		}
		fmt.Println(endpoints)
	},
}

func init() {
	DeployCmd.AddCommand(listCmd)
}
