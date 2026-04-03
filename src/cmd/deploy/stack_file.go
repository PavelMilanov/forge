package deploy

import (
	"fmt"
	"os"

	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var (
	stackFileName string
	stackFilePath string
	stackFileMode string
)

var stackFileCmd = &cobra.Command{
	Use:   "file [endpoint]",
	Short: "Создать или обновить стек из локального файла",
	Long:  "Читает указанный локальный файл и создает/обновляет стек на выбранном endpoint-е согласно --mode.",
	Example: `Создать или обновить:
forge deploy stack file my-endpoint -n my-stack -f ./docker-stack.yml

Явно обновить существующий:
forge deploy stack file my-endpoint -n my-stack -f ./docker-stack.yml --mode update`,
	ValidArgsFunction: endpointCompletion,
	Args:              endpointArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if stackFileMode != "create" && stackFileMode != "update" && stackFileMode != "upsert" {
			errors.DeployErrors(fmt.Errorf("unknown mode %q (allowed: create, update, upsert)", stackFileMode))
		}

		content, err := os.ReadFile(stackFilePath)
		if err != nil {
			errors.DeployErrors(err)
		}

		cfg, endpointID := preparePortainerAndEndpoint(args[0])
		result, err := cfg.DeployStackFromContent(endpointID, stackFileName, string(content), api.StackMode(stackFileMode), true, true)
		if err != nil {
			errors.DeployErrors(err)
		}
		switch result {
		case "created":
			fmt.Printf("Stack %q created on endpoint %q\n", stackFileName, args[0])
		case "updated":
			fmt.Printf("Stack %q updated on endpoint %q\n", stackFileName, args[0])
		default:
			fmt.Printf("Stack %q deployed on endpoint %q\n", stackFileName, args[0])
		}
	},
}

// init регистрирует подкоманду deploy stack file и ее флаги.
func init() {
	stackCmd.AddCommand(stackFileCmd)
	stackFileCmd.Flags().StringVarP(&stackFileName, "name", "n", "", "имя стека")
	stackFileCmd.Flags().StringVarP(&stackFilePath, "file", "f", "", "локальный путь к docker stack файлу")
	stackFileCmd.Flags().StringVar(&stackFileMode, "mode", "upsert", "режим: create | update | upsert")
	stackFileCmd.MarkFlagRequired("name")
	stackFileCmd.MarkFlagRequired("file")
	stackFileCmd.RegisterFlagCompletionFunc("mode", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"create", "update", "upsert"}, cobra.ShellCompDirectiveNoFileComp
	})
}
