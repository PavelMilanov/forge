package env

import (
	"context"
	"fmt"
	"strings"

	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var params []string

var setCmd = &cobra.Command{
	Use:   "set [project]",
	Short: "Обновить параметры deploy-модели проекта",
	Long:  "Загружает текущий секрет проекта из Vault и обновляет ключи в environments.<env>.data по переданным параметрам key=value.",
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
		state, err := loadProjectState(ctx, args[0])
		if err != nil {
			errors.VaultErrors(err)
		}
		ref, ok := getEnvironment(state, environmentName)
		if !ok {
			errors.VaultErrors(fmt.Errorf("environment %q not found for project %q", environmentName, args[0]))
		}
		for _, param := range params {
			key, value, err := parseParam(param)
			if err != nil {
				errors.SpecErrors(err)
			}
			if key == "placement" || strings.HasPrefix(key, "placement.") {
				placementPath := strings.TrimPrefix(key, "placement.")
				if placementPath == "" {
					errors.SpecErrors(fmt.Errorf("placement root is not assignable, use placement.<field>=<value>"))
				}
				setPath(ref.Placement, placementPath, value)
				continue
			}
			setPath(ref.Data, key, value)
		}
		state.Environments[environmentName] = ref
		if err := saveProjectState(ctx, args[0], state); err != nil {
			errors.VaultErrors(err)
		}
		out, err := marshalPrettyJSON(map[string]any{
			"placement": ref.Placement,
			"data":      ref.Data,
		})
		if err != nil {
			errors.VaultErrors(err)
		}
		fmt.Println(out)
	},
}

// init регистрирует подкоманду env set и флаг списка параметров обновления.
func init() {
	EnvCmd.AddCommand(setCmd)
	setCmd.Flags().StringSliceVarP(&params, "param", "p", []string{}, "project parameter")
	setCmd.Flags().StringVarP(&environmentName, "env", "e", "", "environment name: stage | prod")
	setCmd.MarkFlagRequired("env")
}
