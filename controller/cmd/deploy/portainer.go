package deploy

import (
	"fmt"

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
	defaultFlags(portainerCmd)
}
