package deploy

import "github.com/spf13/cobra"

var stackCmd = &cobra.Command{
	Use:       "stack [command]",
	Short:     "Операции со стеками на endpoint-е",
	Long:      "Группа команд stack управляет стеком через Portainer API: разворачивание из файла, создание из custom template и обновление существующего стека.",
	Example:   "forge deploy stack file my-endpoint -n my-stack -f ./docker-stack.yml --mode upsert",
	ValidArgs: []string{"file", "template", "refresh"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// init регистрирует группу подкоманд deploy stack.
func init() {
	DeployCmd.AddCommand(stackCmd)
}
