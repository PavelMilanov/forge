package deploy

import (
	"github.com/PavelMilanov/forge/agent"
	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Развертывание стеков",
	Long:  "Позволяет развертывать стеки на основе шаблонов или обновлять существующие стеки.",
	Example: `Деплой нового стека из шаблона:
forge deploy stack -n my-stack -t my-template

Обновление стека:
forge deploy stack -n my-stack
`,
	ValidArgs: []string{"stack"},
	Args:      cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := agent.NewAgent()
		if deployTemplate != "" {
			if err := cfg.CreateStack(deployStack, deployTemplate); err != nil {
				errors.DeployErrors(err)
			}
		} else {
			if err := cfg.DeployStack(deployStack); err != nil {
				errors.DeployErrors(err)
			}
		}

	},
}

func init() {
	DeployCmd.AddCommand(stackCmd)
	stackCmd.Flags().StringVarP(&deployStack, "name", "n", "", "имя стека")
	stackCmd.Flags().StringVarP(&deployTemplate, "template", "t", "", "шаблон стека")
	stackCmd.MarkFlagRequired("name")
}
