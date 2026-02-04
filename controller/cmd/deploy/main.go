package deploy

import (
	"fmt"
	"strings"

	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
	"github.com/spf13/cobra"
)

var (
	portainerStack string
	portainer      *api.Portainer
)

var DeployCmd = &cobra.Command{
	Use:       "deploy [command]",
	Short:     "Manage deployment",
	Example:   "forge deploy",
	ValidArgs: []string{"portainer"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	env, err := config.NewEnv(config.FORGE_PATH, config.FORGE_FILE)
	if err != nil {
		fmt.Println(err)
	}
	portainer, err = api.NewPortainer(env)
	if err != nil {
		fmt.Println(err)
	}

}

func defaultFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&portainerStack, "stack", "s", "", "portainer stack name")
	cmd.MarkFlagRequired("stack")

	cmd.RegisterFlagCompletionFunc("stack", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
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
