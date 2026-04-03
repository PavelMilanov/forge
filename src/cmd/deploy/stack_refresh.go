package deploy

import (
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var stackRefreshName string

var stackRefreshCmd = &cobra.Command{
	Use:   "refresh [endpoint]",
	Short: "Пере-применить текущий контент существующего стека",
	Long:  "Получает текущий StackFileContent существующего стека из Portainer и отправляет его обратно как update (prune/pullImage=true). Команда не принимает локальный файл и не использует template.",
	Example: `Пере-применить текущую конфигурацию стека:
forge deploy stack refresh my-endpoint -n my-stack`,
	ValidArgsFunction: endpointCompletion,
	Args:              endpointArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, endpointID := preparePortainerAndEndpoint(args[0])
		if err := cfg.RefreshStack(endpointID, stackRefreshName, true, true); err != nil {
			errors.DeployErrors(err)
		}

		fmt.Printf("Stack %q refreshed on endpoint %q\n", stackRefreshName, args[0])
	},
}

// init регистрирует подкоманду deploy stack refresh и ее флаг имени стека.
func init() {
	stackCmd.AddCommand(stackRefreshCmd)
	stackRefreshCmd.Flags().StringVarP(&stackRefreshName, "name", "n", "", "имя существующего стека")
	stackRefreshCmd.MarkFlagRequired("name")
}
