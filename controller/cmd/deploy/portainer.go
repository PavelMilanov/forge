package deploy

import (
	"fmt"

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
			fmt.Println(err)
		}
		for _, stack := range stacks {
			if stack.Name == portainerStack {
				project, err := portainer.GetStackFile(stack)
				if err != nil {
					fmt.Println(err)
				}
				resp, err := portainer.UpdateStack(project)
				if err != nil {
					fmt.Println(err)
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
