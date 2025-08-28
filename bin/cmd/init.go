package cmd

import (
	"context"

	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/spec"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Project initialization",
	Example: `
forge init -f path/to/configFile.yaml -m compose -a dev
`,

	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		project, err := spec.NewSpec(projectMode)
		if err != nil {
			errors.SpecErrors(err)
		}
		project.Init()
		_, err = vault.API.Get(ctx, projectAlias) // ищет указанный проект в хранилище, если не найден, то создаем новый
		if err != nil {
			_, err = vault.API.Put(ctx, projectAlias, map[string]any{"deploy": project, "mode": projectMode})
			if err != nil {
				errors.VaultErrors(err)
			}
			text := fmt.Sprintf("The project %s initialization was successful\nSee %s", projectAlias, vault.ENV.Vault.Url)
			fmt.Println(text)
		}
		fmt.Println("The project already initialized")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	defaultFlags(initCmd)
}
