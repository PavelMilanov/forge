package template

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Показать список доступных шаблонов",
	Long:    "Читает содержимое каталога шаблонов и выводит имена файлов, которые можно использовать при инициализации проекта (`forge env init -t ...`).",
	Example: "forge templates list",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appTemplate.GetTemplates()
	},
}

// init регистрирует подкоманду templates list.
func init() {
	TmpCmd.AddCommand(listCmd)
}
