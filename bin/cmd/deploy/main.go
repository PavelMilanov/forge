package deploy

import (
	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var (
	projectTemplate string
	projectMode     string
	projectAlias    string
	vault           *api.VaultAPI
	template        string
)

var DeployCmd = &cobra.Command{
	Use:   "deploy [command]",
	Short: "Deploy project",
	// Example: "forge deploy",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// ctx := context.Background()
		// secrets, err := vault.API.Get(ctx, args[0])
		// if err != nil {
		// 	errors.VaultErrors(err)
		// }
		// value, exists := secrets.Data["host"]
		// if !exists {
		// 	errors.VaultErrors(fmt.Errorf("host not found"))
		// }
		// host := remote.NewHost()
		// host.Parse(value.(map[string]any))
		// ssh, err := remote.NewSSH(vault.ENV, host.Addr)
		// if err != nil {
		// 	errors.RemoteErrors(err)
		// }
		// defer ssh.Close()
		// // localFile := filepath.Join(config.CONFIG_PATH, "config.yaml")
		// remoteFile := filepath.Join(host.Path, "config.yaml")
		// if err := ssh.Upload(file, remoteFile); err != nil {
		// 	errors.RemoteErrors(err)
		// }
	},
}

func init() {
	var err error
	vault, err = api.NewVaultClient()
	if err != nil {
		errors.VaultErrors(err)
	}
	vault.Set()
	if err := vault.RenewToken(); err != nil {
		errors.VaultErrors(err)
	}

}

func defaultFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&projectTemplate, "file", "f", "", "path to project/to/template.yml")
	cmd.Flags().StringVarP(&projectMode, "mode", "m", "compose", "project mode: compose | swarm | kubernetes")
	cmd.Flags().StringVarP(&projectAlias, "alias", "a", "", "unique alias for the project")
	// cmd.Flags().StringVarP(&hostPath, "path", "p", "/var/app", "path to remote host project directory")
	// cmd.Flags().StringVarP(&hostAddr, "remote", "r", "", "remote host address")
	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("alias")
	// cmd.MarkFlagRequired("remote")
}
