package deploy

import (
	"fmt"

	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
	"github.com/spf13/cobra"
)

var (
	portainerStack string
	portainer      *api.Portainer
)

var DeployCmd = &cobra.Command{
	Use:     "deploy",
	Short:   "Manage deployment",
	Example: "forge deploy",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	env, err := config.NewEnv(config.FORGE_PATH, config.FORGE_FILE)
	if err != nil {
		fmt.Println(err)
	}
	portainer, err = api.NewPortainer(env)
	if err != nil {
		fmt.Println(err)
	}

}

func defaultFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&portainerStack, "stack", "s", "", "portainer stack name")
	cmd.MarkFlagRequired("stack")
}
