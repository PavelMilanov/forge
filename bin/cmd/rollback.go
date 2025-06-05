/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/PavelMilanov/forge/utils"
	"github.com/spf13/cobra"
)

// rollbackCmd represents the rollback command
var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rolled back the service to the specified version",
	Long: `Rolled back the service to the specified version and generating the docker configuration file.
For example:

forge -f docker-compose.yml -s alpine rollback 2 backend
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
		toVersion, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = vault.KV.Rollback(ctx, args[1], toVersion)
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
		text := fmt.Sprintf("Project file %s rolled back.", file)
		fmt.Println(text)
	},
}

func init() {
	rootCmd.AddCommand(rollbackCmd)
	rollbackCmd.PersistentFlags().StringVarP(&dockerFile, "file", "f", "", "forge -f <docker-compose.yml|docker-stack.yml>")
	rollbackCmd.PersistentFlags().StringVarP(&dockerService, "service", "s", "", "forge -s <service_name>")
	rollbackCmd.PersistentFlags().StringVarP(&dockerEnv, "env", "e", "default", "forge -e <env_name>")
	rollbackCmd.MarkPersistentFlagRequired("file")
	rollbackCmd.MarkPersistentFlagRequired("service")
}
