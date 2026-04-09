package stack

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PavelMilanov/forge/errors"
	istack "github.com/PavelMilanov/forge/internal/stack"
	stackdeploy "github.com/PavelMilanov/forge/internal/stack/apply"
	"github.com/PavelMilanov/forge/internal/stack/noop"
	"github.com/spf13/cobra"
)

var (
	applyRuntime string
	applyTarget  string
	applyName    string
	applyFile    string
	applyPrune   bool
	applyPull    bool
	applyGroups  []string
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Применить изменения",
	Long:  "Применяет подготовленный манифест к целевой среде.",
	Run: func(cmd *cobra.Command, args []string) {
		runtime := istack.Runtime(strings.ToLower(applyRuntime))

		svc := stackdeploy.NewService(noop.TargetResolver{}, noop.WorkloadRepository{}, noop.WorkloadApplier{}, noop.AccessManager{})
		result, err := svc.Deploy(context.Background(), istack.ApplyInput{
			Runtime:      runtime,
			TargetName:   applyTarget,
			WorkloadName: applyName,
			ManifestFile: applyFile,
			Prune:        applyPrune,
			PullImage:    applyPull,
			AccessGroups: applyGroups,
		})
		if err != nil {
			errors.DeployErrors(err)
		}

		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			errors.DeployErrors(err)
		}
		fmt.Println(string(data))
	},
}

func init() {
	StackCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringVar(&applyRuntime, "runtime", string(istack.RuntimeSwarm), "runtime: compose | swarm | kubernetes")
	applyCmd.Flags().StringVarP(&applyTarget, "target", "e", "", "имя целевой среды (endpoint/context/cluster)")
	applyCmd.Flags().StringVarP(&applyName, "name", "n", "", "имя workload")
	applyCmd.Flags().StringVarP(&applyFile, "file", "f", "", "путь к подготовленному манифесту")
	applyCmd.Flags().BoolVar(&applyPrune, "prune", true, "runtime-specific prune при update")
	applyCmd.Flags().BoolVar(&applyPull, "pull-image", false, "runtime-specific pull image при update")
	applyCmd.Flags().StringSliceVar(&applyGroups, "access-groups", nil, "группы/команды для прав управления workload")
	applyCmd.MarkFlagRequired("target")
	applyCmd.MarkFlagRequired("name")
	applyCmd.MarkFlagRequired("file")
	applyCmd.RegisterFlagCompletionFunc("runtime", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{string(istack.RuntimeCompose), string(istack.RuntimeSwarm), string(istack.RuntimeKubernetes)}, cobra.ShellCompDirectiveNoFileComp
	})
}
