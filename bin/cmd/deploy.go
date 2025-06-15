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
	Example: `
forge deploy -f test.docker-compose.yml -a stage
<test.docker-compose.yml>
services:
  alpine:
    image: alpine:{{ tag "alpine" }}
    container_name: alpine
    restart: unless-stopped

<stage-stack.yml>
services:
  alpine:
    image: alpine:latest
    container_name: alpine
    restart: unless-stopped
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		secrets, err := vault.KV.Get(ctx, dockerAlias)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tags := map[string]string{}
		for key, value := range secrets.Data {
			tags[key] = value.(string)
		}
		file, err := utils.GenerateAppConfig(dockerFile, dockerAlias, tags)
		if err != nil {
			fmt.Println("Error generating config:", err)
			os.Exit(1)
		}
		text := fmt.Sprintf("Project file %s generated", file)
		fmt.Println(text)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	addDefaultFlags(deployCmd)
}
