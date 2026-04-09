package env

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [project]",
	Short: "Добавить окружение в проект",
	Long:  "Добавляет новое окружение в проектный секрет Vault без изменения шаблона. Для окружения создаются блоки placement и data.",
	Example: `Добавить окружение stage:
forge env add api -e stage

Добавить окружение prod:
forge env add api -e prod`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		state, err := loadProjectState(ctx, args[0])
		if err != nil {
			errors.VaultErrors(err)
		}
		if err := fillEnvironmentDataFromTemplate(&state, environmentName); err != nil {
			errors.SpecErrors(err)
		}
		if err := saveProjectState(ctx, args[0], state); err != nil {
			errors.VaultErrors(err)
		}
		out, err := marshalPrettyJSON(map[string]any{
			"template": state.Template,
			environmentName: map[string]any{
				"placement": state.Environments[environmentName].Placement,
				"data":      state.Environments[environmentName].Data,
			},
		})
		if err != nil {
			errors.VaultErrors(err)
		}
		fmt.Println(out)
	},
}

func init() {
	EnvCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&environmentName, "env", "e", "", "environment name: stage | prod")
	addCmd.MarkFlagRequired("env")
	addCmd.RegisterFlagCompletionFunc("env", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"stage", "prod"}, cobra.ShellCompDirectiveNoFileComp
	})
}
