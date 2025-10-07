package deploy

import (
	"context"
	"os"

	"fmt"

	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/spec"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init [FLAGS]",
	Short:   "Project initialization",
	Example: "forge init -f path/to/template.yaml -m compose -a <string>",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		project, err := spec.NewSpec(projectMode)
		if err != nil {
			errors.SpecErrors(err)
		}
		project.Init()
		// host := remote.Host{Addr: hostAddr, Path: hostPath}
		_, err = vault.API.Get(ctx, projectAlias) // ищет указанный проект в хранилище, если не найден, то создаем новый
		if err != nil {
			_, err = vault.API.Put(ctx, projectAlias, map[string]any{
				"deploy": project,
				"mode":   projectMode,
				// "host":   host,
			})
			if err != nil {
				errors.VaultErrors(err)
			}
			text := fmt.Sprintf("The project %s initialization was successful\nSee %s", projectAlias, vault.ENV.Vault.Url)
			fmt.Println(text)
			os.Exit(0)
		}
		fmt.Println("The project already initialized")
	},
}

func init() {
	DeployCmd.AddCommand(initCmd)
	defaultFlags(initCmd)
}
