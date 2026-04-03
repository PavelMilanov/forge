package env

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/spec"
	"github.com/spf13/cobra"
)

var params []string

var setCmd = &cobra.Command{
	Use:   "set [project]",
	Short: "Обновить параметры deploy-модели проекта",
	Long:  "Загружает текущий секрет проекта из Vault, валидирует переданные параметры относительно режима проекта (compose/swarm) и сохраняет обновленные значения как новую версию секрета.",
	Example: `Обновить тег образа:
forge env set dev -p tag=1.2.3

Обновить образ и количество реплик (для swarm):
forge env set stage -p image=registry.local/app -p replicas=3`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(params) == 0 {
			errors.SpecErrors(fmt.Errorf("no parameters detected"))
		}
		ctx := context.Background()
		secrets, err := vault.API.Get(ctx, args[0])
		if err != nil {
			errors.VaultErrors(err)
		}
		data := secrets.Data
		project, err := spec.NewSpec(data["mode"].(string))
		if err != nil {
			errors.SpecErrors(err)
		}
		project.Parse(data)
		if err := project.Update(params); err != nil {
			errors.SpecErrors(err)
		}
		_, err = vault.API.Patch(ctx, args[0], map[string]any{"deploy": project})
		if err != nil {
			errors.VaultErrors(err)
		}
		project.Print()
	},
}

// init регистрирует подкоманду env set и флаг списка параметров обновления.
func init() {
	EnvCmd.AddCommand(setCmd)
	setCmd.Flags().StringSliceVarP(&params, "param", "p", []string{}, "project parameter")
}
