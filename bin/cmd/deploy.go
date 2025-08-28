package cmd

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/spec"
	"github.com/spf13/cobra"
)

var file string

var deployCmd = &cobra.Command{
	Use:     "deploy [OPTIONS] [FLAGS]",
	Short:   "Generating a project configuration file",
	Example: "forge deploy dev -f test.docker-compose.yml",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		secrets, err := vault.API.Get(ctx, args[0])
		if err != nil {
			errors.VaultErrors(err)
		}
		value, exists := secrets.Data["deploy"]
		if !exists {
			errors.VaultErrors(fmt.Errorf("value not found"))
		}
		project, err := spec.NewSpec(secrets.Data["mode"].(string))
		if err != nil {
			errors.SpecErrors(err)
		}
		project.Parse(value.(map[string]any))
		config, err := project.Generate(file, args[0])
		if err != nil {
			errors.SpecErrors(err)
		}
		fmt.Printf("File generated: %s\n", config)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().StringVarP(&file, "file", "f", "", "path/to/file")
}
