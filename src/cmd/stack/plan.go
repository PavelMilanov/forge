package stack

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PavelMilanov/forge/errors"
	istack "github.com/PavelMilanov/forge/internal/stack"
	"github.com/PavelMilanov/forge/internal/stack/noop"
	prepare "github.com/PavelMilanov/forge/internal/stack/plan"
	"github.com/spf13/cobra"
)

var (
	planProject string
	planTpl     string
	planOutput  string
	planRuntime string
	planTarget  string
	planName    string
)

type planResult struct {
	Prepare istack.PlanResult `json:"prepare"`
	Intent  struct {
		Runtime      istack.Runtime `json:"runtime"`
		TargetName   string         `json:"targetName,omitempty"`
		WorkloadName string         `json:"workloadName,omitempty"`
		ManifestFile string         `json:"manifestFile"`
	} `json:"intent"`
}

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Сформировать план изменений (без применения)",
	Long:  "Подготавливает манифест и выводит план применения без выполнения deploy-операций.",
	Run: func(cmd *cobra.Command, args []string) {
		runtime := istack.Runtime(strings.ToLower(planRuntime))

		svc := prepare.NewService(noop.PrepareSource{}, noop.Renderer{}, noop.Validator{}, noop.Writer{})
		result, err := svc.Prepare(context.Background(), istack.PlanInput{
			Project:    planProject,
			Template:   planTpl,
			Runtime:    runtime,
			OutputPath: planOutput,
		})
		if err != nil {
			errors.DeployErrors(err)
		}

		out := planResult{Prepare: result}
		out.Intent.Runtime = runtime
		out.Intent.TargetName = planTarget
		out.Intent.WorkloadName = planName
		out.Intent.ManifestFile = planOutput

		data, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			errors.DeployErrors(err)
		}
		fmt.Println(string(data))
	},
}

func init() {
	StackCmd.AddCommand(planCmd)
	planCmd.Flags().StringVarP(&planProject, "project", "p", "", "проект в Vault")
	planCmd.Flags().StringVarP(&planTpl, "template", "t", "", "опциональный шаблон")
	planCmd.Flags().StringVarP(&planOutput, "output", "o", "", "куда сохранить итоговый манифест")
	planCmd.Flags().StringVar(&planRuntime, "runtime", string(istack.RuntimeSwarm), "runtime: compose | swarm | kubernetes")
	planCmd.Flags().StringVarP(&planTarget, "target", "e", "", "целевая среда для будущего apply")
	planCmd.Flags().StringVarP(&planName, "name", "n", "", "имя workload для будущего apply")
	planCmd.MarkFlagRequired("project")
	planCmd.MarkFlagRequired("output")
	planCmd.RegisterFlagCompletionFunc("runtime", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{string(istack.RuntimeCompose), string(istack.RuntimeSwarm), string(istack.RuntimeKubernetes)}, cobra.ShellCompDirectiveNoFileComp
	})
}
