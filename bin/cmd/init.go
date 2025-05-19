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
		stack, _ := docker.NewStack(dockerFile)
		data := map[string]interface{}{}
		for _, service := range stack.App.Services {
			data[service.Name] = "undefined"
		}
		_, err := vault.Put(ctx, args[0], data)
		if err != nil {
			os.Exit(1)
		}
		text := fmt.Sprintf("The project %s initialization was successful.", args[0])
		fmt.Println(text)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
