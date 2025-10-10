package server

import (
	"os"

	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/errors"
	"github.com/spf13/cobra"
)

var (
	env *config.Env
)

var ServerCmd = &cobra.Command{
	Use:     "server [COMMAND] [OPTIONS] [FLAGS]",
	Short:   "server project",
	Example: "server",
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
	env, err = config.NewEnv(config.CONFIG_PATH, config.CONFIG_FILE)
	if err != nil {
		errors.RemoteErrors(err)
		os.Exit(1)
	}

}
