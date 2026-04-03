package deploy

import (
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var (
	stackTemplateStackName string
	stackTemplateName      string
)

var stackTemplateCmd = &cobra.Command{
	Use:   "template [endpoint]",
	Short: "Создать стек из custom template Portainer",
	Long:  "Берет custom template в Portainer, рендерит его и создает новый стек на выбранном endpoint-е.",
	Example: `Создать стек из template:
forge deploy stack template my-endpoint -n my-stack -t base`,
	ValidArgsFunction: endpointCompletion,
	Args:              endpointArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, endpointID := preparePortainerAndEndpoint(args[0])

		content, err := buildTemplateContent(cfg, stackTemplateName, stackTemplateStackName)
		if err != nil {
			errors.DeployErrors(err)
		}

		if err := cfg.CreateStack(endpointID, stackTemplateStackName, content); err != nil {
			errors.DeployErrors(err)
		}
		fmt.Printf("Stack %q created on endpoint %q from template %q\n", stackTemplateStackName, args[0], stackTemplateName)
	},
}

// init регистрирует подкоманду deploy stack template и ее флаги.
func init() {
	stackCmd.AddCommand(stackTemplateCmd)
	stackTemplateCmd.Flags().StringVarP(&stackTemplateStackName, "name", "n", "", "имя стека")
	stackTemplateCmd.Flags().StringVarP(&stackTemplateName, "template", "t", "", "имя custom template в Portainer")
	stackTemplateCmd.MarkFlagRequired("name")
	stackTemplateCmd.MarkFlagRequired("template")
}
