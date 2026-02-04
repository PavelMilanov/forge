package deploy

import (
	"fmt"
	"strings"

	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var portainerCmd = &cobra.Command{
	Use:     "portainer",
	Short:   "Portainer deployment",
	Example: "forge deploy portainer",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		stacks, err := portainer.GetStacks()
		if err != nil {
			errors.DeployErrors(err)
		}
		for _, stack := range stacks {
			if stack.Name == portainerStack {
				project, err := portainer.GetStackFile(stack)
				if err != nil {
					errors.DeployErrors(err)
				}
				resp, err := portainer.UpdateStack(project)
				if err != nil {
					errors.DeployErrors(err)
				}
				fmt.Println(resp)
				return
			}
		}
		fmt.Printf("stack %s не найден\n", portainerStack)
	},
}

func init() {
	DeployCmd.AddCommand(portainerCmd)
	portainerCmd.Flags().StringVarP(&portainerStack, "stack", "s", "", "portainer stack name")
	portainerCmd.MarkFlagRequired("stack")

	portainerCmd.RegisterFlagCompletionFunc("stack", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if portainer == nil {
			if config.AppConfig == nil {
				cfg, err := config.NewEnv(config.FORGE_PATH, config.FORGE_FILE)
				if err != nil {
					return nil, cobra.ShellCompDirectiveError
				}
				config.AppConfig = cfg
			}
			p, err := api.NewPortainer(config.AppConfig)
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			portainer = p
		}
		stacks, err := portainer.GetStacks()
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}

		var completions []string
		for _, stack := range stacks {
			if strings.HasPrefix(stack.Name, toComplete) {
				completions = append(completions, stack.Name)
			}
		}
		return completions, cobra.ShellCompDirectiveNoFileComp
	})
}
