package env

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project]",
	Short: "Создать начальный секрет проекта в Vault",
	Long:  "Инициализирует проект в Vault: сохраняет шаблон и создает окружение с блоками placement и data. Если проект уже существует, команда дополняет структуру недостающим окружением.",
	Example: `Инициализация проекта для stage:
forge env init my-app -e stage -t my-template.yaml

Инициализация проекта для prod:
forge env init admin -e prod -t stack.yml`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		state, err := loadProjectState(ctx, args[0])
		if err != nil {
			if containsSecretNotFound(err) {
				state = projectState{
					Template:     projectTemplate,
					Environments: map[string]environmentRef{},
				}
			} else {
				errors.VaultErrors(err)
			}
		}
		if state.Template == "" {
			state.Template = projectTemplate
		}
		if err := fillEnvironmentDataFromTemplate(&state, environmentName); err != nil {
			errors.SpecErrors(err)
		}
		if err := saveProjectState(ctx, args[0], state); err != nil {
			errors.VaultErrors(err)
		}
		text := fmt.Sprintf("The project %s initialization was successful\nSee %s", args[0], vault.ENV.Vault.Url)
		fmt.Println(text)
	},
}

// init регистрирует подкоманду env init и настраивает обязательные флаги.
func init() {
	EnvCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&projectTemplate, "template", "t", "", "path to project/to/template.yml")
	initCmd.Flags().StringVarP(&environmentName, "env", "e", "", "environment name: stage | prod")
	initCmd.MarkFlagRequired("template")
	initCmd.MarkFlagRequired("env")
	initCmd.RegisterFlagCompletionFunc("env", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"stage", "prod"}, cobra.ShellCompDirectiveNoFileComp
	})
}
