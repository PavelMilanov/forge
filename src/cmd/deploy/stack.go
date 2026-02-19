package deploy

import (
	"github.com/PavelMilanov/forge/agent"
	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var stackCmd = &cobra.Command{
	Use:   "stack [endpoint] [flags]",
	Short: "Развертывание стеков",
	Long:  "Позволяет развертывать стеки на основе шаблонов или обновлять существующие стеки.",
	Example: `Деплой нового стека из шаблона:
forge deploy stack my-endpoint -n my-stack -t my-template

Обновление стека:
forge deploy stack my-endpoint -n my-stack
`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if err := loadEndpointAliases(); err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return endpointAliases, cobra.ShellCompDirectiveNoFileComp
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := loadEndpointAliases(); err != nil {
			return err
		}
		for _, name := range endpoints {
			endpointAliases = append(endpointAliases, name)
		}
		cmd.ValidArgs = endpointAliases
		return cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		for idx, name := range endpoints {
			if args[0] == name {
				deplyEndpoint = idx
				return
			}
		}
		cfg := agent.NewAgent()
		if deployTemplate != "" {
			if err := cfg.CreateStack(deplyEndpoint, deployStack, deployTemplate); err != nil {
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

func loadEndpointAliases() error {
	// if len(endpointAliases) > 0 {
	// 	return nil // Уже загружено
	// }
	var err error
	cfg := agent.NewAgent()
	endpoints, err = cfg.ListEndpoints()
	if err != nil {
		return err
	}
	// for _, name := range endpoints {
	// 	endpointAliases = append(endpointAliases, name)
	// }
	return nil
}
