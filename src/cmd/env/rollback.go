package env

import (
	"context"

	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/spec"
	"github.com/spf13/cobra"
)

var versionNum int

var rollbackCmd = &cobra.Command{
	Use:   "rollback [project] [flag]",
	Short: "Откатить секрет проекта к выбранной версии",
	Long:  "Выполняет rollback данных проекта в Vault к указанной версии и выводит восстановленные значения deploy-модели. Команда не применяет изменения в инфраструктуре автоматически и не выполняет деплой.",
	Example: `Откатить проект dev к версии 1:
forge env rollback dev -v 1

Перед откатом посмотреть доступные версии:
forge env versions dev`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		secrets, err := vault.API.Rollback(ctx, args[0], versionNum)
		if err != nil {
			errors.VaultErrors(err)
		}
		data := secrets.Data
		project, err := spec.NewSpec(data["mode"].(string))
		project.Parse(data)
		if err != nil {
			errors.VaultErrors(err)
		}
		project.Print()
	},
}

// init регистрирует подкоманду env rollback и обязательный флаг версии.
func init() {
	EnvCmd.AddCommand(rollbackCmd)
	rollbackCmd.Flags().IntVarP(&versionNum, "version", "v", 0, "number version")
	rollbackCmd.MarkFlagRequired("version")
}
