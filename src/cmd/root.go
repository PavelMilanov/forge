package cmd

import (
	"os"

	"github.com/PavelMilanov/forge/cmd/deploy"
	"github.com/PavelMilanov/forge/cmd/env"
	"github.com/PavelMilanov/forge/cmd/template"
	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "forge",
	Short:   "CLI для управления переменными деплоя, шаблонами и развертыванием инфраструктуры",
	Long:    "forge объединяет работу с Vault, шаблонами конфигурации и агентом деплоя. Используйте подкоманды env, tmpl и deploy для управления жизненным циклом конфигурации сервисов.",
	Example: "forge env init admin-stage -t stack.yml -m swarm",
	Version: config.VERSION,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute запускает корневую CLI-команду и завершает процесс с кодом 1 при ошибке.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init инициализирует конфигурацию приложения и подключает корневые подкоманды CLI.
func init() {
	cfg, err := config.NewEnv(config.FORGE_PATH, config.FORGE_FILE)
	if err != nil {
		errors.ForgeErrors(err)
	}
	config.NewAppConfig(cfg)
	rootCmd.AddCommand(env.EnvCmd)
	rootCmd.AddCommand(template.TmpCmd)
	rootCmd.AddCommand(deploy.DeployCmd)
}
