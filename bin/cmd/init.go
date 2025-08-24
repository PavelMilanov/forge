package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Project initialization",
	Example: `
forge init -f docker/test/docker-compose.yaml -a backend
`,

	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// ctx := context.Background()
		// stack, err := docker.NewStack(dockerFile, dockerAlias)
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
		// data := map[string]interface{}{}
		// for _, service := range stack.App.Services {
		// 	data[service.Name] = strings.Split(service.Image, ":")[1]
		// }
		// _, err = vault.KV.Get(ctx, dockerAlias)
		// if err != nil {
		// 	_, err = vault.KV.Put(ctx, dockerAlias, data)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		os.Exit(1)
		// 	}
		// 	text := fmt.Sprintf("The project %s initialization was successful\nSee %s", dockerAlias, vault.ENV.Vault.Url)
		// 	fmt.Println(text)
		// 	os.Exit(0)
		// }
		fmt.Println("The project already initialized")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	addDefaultFlags(initCmd)
}
