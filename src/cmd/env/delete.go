package env

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete [project]",
	Aliases: []string{"rm"},
	Short:   "Удалить окружение из проекта",
	Long:    "Удаляет окружение из блока environments проектного секрета Vault.",
	Example: `Удалить окружение stage:
forge env delete api -e stage

Удалить окружение prod:
forge env rm api -e prod`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		state, err := loadProjectState(ctx, args[0])
		if err != nil {
			errors.VaultErrors(err)
		}
		if _, ok := state.Environments[environmentName]; !ok {
			errors.VaultErrors(fmt.Errorf("environment %q not found for project %q", environmentName, args[0]))
		}
		delete(state.Environments, environmentName)
		if err := saveProjectState(ctx, args[0], state); err != nil {
			errors.VaultErrors(err)
		}
		out, err := marshalPrettyJSON(map[string]any{
			"project":             args[0],
			"deleted_environment": environmentName,
		})
		if err != nil {
			errors.VaultErrors(err)
		}
		fmt.Println(out)
	},
}

func init() {
	EnvCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&environmentName, "env", "e", "", "environment name: stage | prod")
	deleteCmd.MarkFlagRequired("env")
	deleteCmd.RegisterFlagCompletionFunc("env", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"stage", "prod"}, cobra.ShellCompDirectiveNoFileComp
	})
}
