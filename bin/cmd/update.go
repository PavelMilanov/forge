package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/PavelMilanov/forge/utils"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updating the service to the specified version",
	Long: `Updating the service to the specified version and generating the docker configuration file.
For example:
forge -f docker-compose.yml -s alpine update 3.21 backend
<backend-stack.yml>
services:
  alpine:
    image: alpine:3.21
    container_name: alpine
    restart: unless-stopped
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		data := map[string]interface{}{}
		data[dockerService] = args[0]
		_, err := vault.KV.Patch(ctx, args[1], data)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		secrets, err := vault.KV.Get(ctx, args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tags := map[string]string{}
		for key, value := range secrets.Data {
			tags[key] = value.(string)
		}
		file, err := utils.GenerateAppConfig(dockerFile, args[1], tags)
		if err != nil {
			fmt.Println("Error generating config:", err)
			os.Exit(1)
		}
		text := fmt.Sprintf("Project file updated.\nSee %s", file)
		fmt.Println(text)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().StringVarP(&dockerFile, "file", "f", "", "forge -f <docker-compose.yml|docker-stack.yml>")
	updateCmd.PersistentFlags().StringVarP(&dockerService, "service", "s", "", "forge -s <service_name>")
	updateCmd.MarkPersistentFlagRequired("file")
	updateCmd.MarkPersistentFlagRequired("service")
}
