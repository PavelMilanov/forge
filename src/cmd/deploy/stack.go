package deploy

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var stackCmd = &cobra.Command{
	Use:   "stack [endpoint] [flags]",
	Short: "Развертывание стеков",
	Long:  "Позволяет развертывать стеки на основе шаблонов или обновлять существующие стеки.",
	Example: `Деплой нового стека из шаблона:
forge deploy stack my-endpoint -n my-stack -t my-template

Обновление стека:
forge deploy stack my-endpoint -n my-stack
`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if err := loadEndpointAliases(); err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return endpointAliases, cobra.ShellCompDirectiveNoFileComp
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := loadEndpointAliases(); err != nil {
			return err
		}
		for _, item := range endpoints {
			endpointAliases = append(endpointAliases, item.Name)
		}
		cmd.ValidArgs = endpointAliases
		return cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, item := range endpoints {
			if args[0] == item.Name {
				deplyEndpoint = item.ID
			}
		}
		cfg, err := api.NewPortainer(
			config.AppConfig.Agent.Credentials.Url,
			config.AppConfig.Agent.Credentials.Key)
		if err != nil {
			errors.DeployErrors(err)
		}
		if deployTemplate != "" {
			tmpls, err := cfg.GetTemplates()
			if err != nil {
				errors.DeployErrors(err)
			}
			if len(tmpls) == 0 {
				fmt.Println("no stacks found")
			}
			for _, tmpl := range tmpls {
				if tmpl.TemplateName == deployTemplate {
					proj, err := cfg.GetTemplateFile(tmpl)
					if err != nil {
						errors.DeployErrors(err)
					}
					var buf bytes.Buffer
					data := map[string]string{
						"stand": deployStack,
					}
					funcMap := template.FuncMap{
						"port": func(s string) string {
							return strings.TrimPrefix(s, "f") // Удаляет "f", если она есть в начале
						},
					}
					t := template.Must(template.New("portainer").Option("missingkey=error").Funcs(funcMap).Parse(string(proj.TemplateFile)))
					err = t.Execute(&buf, data)
					if err != nil {
						errors.DeployErrors(err)
					}
					if err := cfg.CreateStack(deplyEndpoint, deployStack+"-"+deployTemplate, buf.String()); err != nil {
						errors.DeployErrors(err)
					}
					fmt.Println("Stack created")
				}
			}
		} else {
			stacks, err := cfg.GetStacks()
			if err != nil {
				errors.DeployErrors(err)
			}
			if len(stacks) == 0 {
				fmt.Println("no stacks found")
				errors.DeployErrors(err)
			}
			for _, stack := range stacks {
				if stack.StackName == deployStack {
					project, err := cfg.GetStackFile(stack)
					if err != nil {
						errors.DeployErrors(err)
					}
					if err := cfg.UpdateStack(project); err != nil {
						errors.DeployErrors(err)
					}
					fmt.Println("Stack updated")
					return
				}
				fmt.Println("no stacks found")
			}
		}
	},
}

func init() {
	DeployCmd.AddCommand(stackCmd)
	stackCmd.Flags().StringVarP(&deployStack, "name", "n", "", "имя стека")
	stackCmd.Flags().StringVarP(&deployTemplate, "template", "t", "", "шаблон стека")
	stackCmd.MarkFlagRequired("name")
}

/*
loadEndpointAliases загружает список доступных окружений из конфигурации агента.

Returns

	error - ошибка.
*/
func loadEndpointAliases() error {
	cfg, err := api.NewPortainer(
		config.AppConfig.Agent.Credentials.Url,
		config.AppConfig.Agent.Credentials.Key)
	data, err := cfg.GetEndpoints()
	if err != nil {
		return err
	}
	endpoints = data
	return nil
}
