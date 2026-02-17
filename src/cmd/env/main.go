package env

import (
	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
	"github.com/spf13/cobra"
)

var (
	projectTemplate string
	projectMode     string
	envConfig       bool
	vault           *api.Vault
)

var EnvCmd = &cobra.Command{
	Use:       "env [command]",
	Short:     "Модуль управления окружением.",
	Long:      "Позволяет управлять окружением, включая инициализацию, настройку, слежение за изменениями и откат версий.",
	Example:   "forge env",
	ValidArgs: []string{"init", "set", "get", "rollback", "versions"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := api.NewVaultClient(config.AppConfig)
		if err != nil {
			return err
		}
		if err := cfg.Login(); err != nil {
			return err
		}
		vault = cfg
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
}
