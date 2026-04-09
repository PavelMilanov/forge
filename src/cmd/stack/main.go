package stack

import "github.com/spf13/cobra"

var StackCmd = &cobra.Command{
	Use:   "stack [command]",
	Short: "Планирование и применение манифестов",
	Long:  "Команды stack реализуют двухэтапный процесс: plan (подготовка/план) и apply (применение).",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgs: []string{
		"plan",
		"apply",
	},
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {}
