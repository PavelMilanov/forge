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
	Short:     "Управление конфигурацией проекта в Vault",
	Long:      "Группа команд env отвечает за полный цикл работы с переменными деплоя: инициализацию проекта, изменение параметров, просмотр текущих значений, просмотр истории версий и откат к выбранной версии секрета.",
	Example:   "forge env set admin-stage -p tag=sha-abc1234",
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

// init объявлен для единообразия структуры пакета env.
func init() {
}
