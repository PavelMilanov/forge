package env

import (
	"context"
	"os"

	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/spec"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Project initialization",
	Example: "forge env init [PROJECT] -t template.yaml -m compose",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		project, err := spec.NewSpec(projectMode)
		if err != nil {
			errors.SpecErrors(err)
		}
		project.Init()
		_, err = vault.API.Get(ctx, args[0]) // ищет указанный проект в хранилище, если не найден, то создаем новый
		if err != nil {
			_, err = vault.API.Put(ctx, args[0], map[string]any{
				"deploy":   project,
				"mode":     projectMode,
				"template": projectTemplate,
			})
			if err != nil {
				errors.VaultErrors(err)
			}
			text := fmt.Sprintf("The project %s initialization was successful\nSee %s", args[0], vault.ENV.Vault.Url)
			fmt.Println(text)
			os.Exit(0)
		}
		fmt.Println("The project already initialized")
	},
}

func init() {
	EnvCmd.AddCommand(initCmd)
	defaultFlags(initCmd)
}
