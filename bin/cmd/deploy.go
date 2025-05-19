package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/PavelMilanov/forge/utils"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Generating a project configuration file",
	Long: `Generate a project configuration file based on monitored service versions.
For example:
forge -f docker/test/docker-compose.test2.yaml deploy backend
<docker-compose.template.yml>
services:
  alpine:
    image: alpine:{{ tag "alpine" }}
    container_name: alpine
    restart: unless-stopped

<docker-compose.yml>
services:
  alpine:
    image: alpine:latest
    container_name: alpine
    restart: unless-stopped
`, Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		secrets, err := vault.Get(ctx, args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tags := map[string]string{}
		for key, value := range secrets.Data {
			tags[key] = value.(string)
		}
		if err := utils.GenerateAppConfig(dockerFile, tags); err != nil {
			fmt.Println("Ошибка генерации конфигурационного файла:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
