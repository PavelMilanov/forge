package cmd

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/spec"
	"github.com/spf13/cobra"
)

var template string

var generateCmd = &cobra.Command{
	Use:     "generate [OPTIONS] [FLAGS]",
	Short:   "Generating a project configuration file",
	Example: "forge generate <project> -f <absolute/path/to/template.yml>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		secrets, err := vault.API.Get(ctx, args[0])
		if err != nil {
			errors.VaultErrors(err)
		}
		value, exists := secrets.Data["deploy"]
		if !exists {
			errors.VaultErrors(fmt.Errorf("deployment not found"))
		}
		project, err := spec.NewSpec(secrets.Data["mode"].(string))
		if err != nil {
			errors.SpecErrors(err)
		}
		project.Parse(value.(map[string]any))
		config, err := project.Generate(template, args[0])
		if err != nil {
			errors.SpecErrors(err)
		}
		fmt.Println(config)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&template, "file", "f", "", "path/to/template.yml")
	generateCmd.MarkFlagRequired("file")
}
