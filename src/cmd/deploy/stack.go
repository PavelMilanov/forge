package deploy

import (
	"github.com/PavelMilanov/forge/agent"
	"github.com/PavelMilanov/forge/errors"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var stackCmd = &cobra.Command{
	Use:       "stack",
	Short:     "Stack deployment",
	Example:   "forge deploy stack",
	ValidArgs: []string{"stack"},
	Args:      cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := agent.NewAgent()
		// if deployTemplate != "" {
		// 	if err := cfg.CreateStack(deployStack, deployTemplate); err != nil {
		// 		errors.DeployErrors(err)
		// 	}
		// } else {
		// 	if err := cfg.DeployStack(deployStack); err != nil {
		// 		errors.DeployErrors(err)
		// 	}
		// }
		s := spinner.New()
		s.Spinner = spinner.Dot
		s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

		m := model{
			spinner:  s,
			cfg:      cfg,
			stack:    deployStack,
			template: deployTemplate,
		}

		p := tea.NewProgram(m)

		if _, err := p.Run(); err != nil {
			errors.DeployErrors(err)
		}
	},
}

func init() {
	DeployCmd.AddCommand(stackCmd)
	stackCmd.Flags().StringVarP(&deployStack, "name", "n", "", "stack name")
	stackCmd.Flags().StringVarP(&deployTemplate, "template", "t", "", "stack template")
	stackCmd.MarkFlagRequired("name")
}
