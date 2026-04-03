package env

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var versionsCmd = &cobra.Command{
	Use:   "versions [project]",
	Short: "Показать историю версий секрета проекта",
	Long:  "Выводит список версий секрета проекта из Vault (номер версии и время создания). Команда полезна перед откатом, когда нужно выбрать корректную версию.",
	Example: `forge env versions my-app

# пример результата:
# Version: 4
# Created: 03/04/2026 15:24:11`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		content, err := vault.API.GetVersionsAsList(ctx, args[0])
		if err != nil {
			errors.VaultErrors(err)
		}
		for _, metadata := range content {
			fmt.Printf(`Version: %d
Created: %s

`, metadata.Version, metadata.CreatedTime.Format("02/01/2006 15:04:05"))
		}
	},
}

// init регистрирует подкоманду env versions.
func init() {
	EnvCmd.AddCommand(versionsCmd)
}
