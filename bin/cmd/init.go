package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/PavelMilanov/forge/docker"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Project initialization",
	Long: `Set of versions of monitored services in a docker-compose file.
For example:
forge -f docker/test/docker-compose.yaml init backend
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		stack, err := docker.NewStack(dockerFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		data := map[string]interface{}{}
		for _, service := range stack.App.Services {
			data[service.Name] = "undefined"
		}

		_, err = vault.Get(ctx, args[0])
		if err != nil {
			_, err = vault.Put(ctx, args[0], data)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			text := fmt.Sprintf("The project %s initialization was successful.", args[0])
			fmt.Println(text)
			os.Exit(0)
		}
		fmt.Println("The project already initialized.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
