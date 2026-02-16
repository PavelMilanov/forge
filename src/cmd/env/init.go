package env

import (
	"context"
	"os"

	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/spec"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project]",
	Short: "Инициализирует проект",
	Long:  "Инициализирует проект, создавая необходимые файлы и структуру для работы с проектом.",
	Example: `Первоначальная инициализация проекта:
forge env init my-app -t my-template.yaml -m compose

Предварительно шаблон должен быть инициализирован.

Модуль проекта по умолчанию - compose. Если не указан, то используется значение по умолчанию.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		project, err := spec.NewSpec(projectMode)
		if err != nil {
			errors.SpecErrors(err)
		}
		project.Init()
		_, err = vault.API.Get(ctx, args[0]) // ищет указанный проект в хранилище, если не найден, то создаем новый
		if err != nil {
			_, err = vault.API.Put(ctx, args[0], map[string]any{
				"deploy":   project,
				"mode":     projectMode,
				"template": projectTemplate,
			})
			if err != nil {
				errors.VaultErrors(err)
			}
			text := fmt.Sprintf("The project %s initialization was successful\nSee %s", args[0], vault.ENV.Vault.Url)
			fmt.Println(text)
			os.Exit(0)
		}
		fmt.Println("The project already initialized")
	},
}

func init() {
	EnvCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&projectTemplate, "template", "t", "", "path to project/to/template.yml")
	initCmd.Flags().StringVarP(&projectMode, "mode", "m", "compose", "project mode: compose | swarm | kubernetes")
	initCmd.MarkFlagRequired("template")
	initCmd.MarkFlagRequired("mode")

	initCmd.RegisterFlagCompletionFunc("mode", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"compose", "swarm", "kubernetes"}, cobra.ShellCompDirectiveNoFileComp
	})
}
