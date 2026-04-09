package env

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

/*
forge env get dev | grep 'tag:' | awk '{print $2}'
*/
var getCmd = &cobra.Command{
	Use:   "get [project]",
	Short: "Показать параметры проекта или сгенерированный конфиг",
	Long:  "Получает секрет проекта из Vault и выводит либо текущие параметры deploy-модели, либо итоговый конфигурационный файл, срендеренный из шаблона (`-c`).",
	Example: `Вывод информации об текущих значениях окружения:
forge env get dev

Вывод полной конфигурации проекта:
forge env get dev -c

Извлечение только tag в shell:
forge env get dev | grep 'tag:' | awk '{print $2}'`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		state, err := loadProjectState(ctx, args[0])
		if err != nil {
			errors.VaultErrors(err)
		}
		ref, ok := getEnvironment(state, environmentName)
		if !ok {
			errors.VaultErrors(fmt.Errorf("environment %q not found for project %q", environmentName, args[0]))
		}
		if envConfig {
			if state.Template == "" {
				errors.VaultErrors(fmt.Errorf("template is empty"))
			}
			std, err := renderTemplate(state.Template, renderContextForEnvironment(ref))
			if err != nil {
				errors.VaultErrors(err)
			}
			fmt.Print(std)
		} else {
			result := map[string]any{
				"placement": ref.Placement,
				"data":      ref.Data,
			}
			out, err := marshalPrettyJSON(result)
			if err != nil {
				errors.VaultErrors(err)
			}
			fmt.Println(out)
		}
	},
}

// init регистрирует подкоманду env get и флаг вывода итогового конфига.
func init() {
	EnvCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&envConfig, "config", "c", false, "project config secret")
	getCmd.Flags().StringVarP(&environmentName, "env", "e", "", "environment name: stage | prod")
	getCmd.MarkFlagRequired("env")
}
