package deploy

import (
	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
	"github.com/spf13/cobra"
)

var (
	portainerStack string
	portainer      *api.Portainer
)

var DeployCmd = &cobra.Command{
	Use:       "deploy [command]",
	Short:     "Manage deployment",
	Example:   "forge deploy",
	ValidArgs: []string{"portainer"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := api.NewPortainer(config.AppConfig)
		if err != nil {
			return err
		}
		portainer = cfg
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
}
