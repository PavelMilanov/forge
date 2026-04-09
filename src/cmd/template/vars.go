package template

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/templatevars"
	"github.com/spf13/cobra"
)

var varsCmd = &cobra.Command{
	Use:   "vars [template]",
	Short: "Извлечь изменяемые переменные Go template из YAML",
	Long:  "Читает YAML-шаблон и извлекает переменные в структуру variables.",
	Example: `forge templates vars api-stack.yml
forge templates vars /var/forge/templates/api-stack.yml`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		candidates := []string{
			args[0],
			filepath.Join(config.TEMPLATE_PATH, args[0]),
		}

		var (
			content []byte
			err     error
		)
		for _, path := range candidates {
			content, err = os.ReadFile(path)
			if err == nil {
				break
			}
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		result, err := templatevars.ExtractFromYAML(content)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(result); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	TmpCmd.AddCommand(varsCmd)
}
