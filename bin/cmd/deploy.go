package cmd

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/PavelMilanov/forge/config"
	"github.com/PavelMilanov/forge/errors"
	"github.com/PavelMilanov/forge/remote"
	"github.com/spf13/cobra"
)

var file string

var deployCmd = &cobra.Command{
	Use:     "deploy [OPTIONS] [FLAGS]",
	Short:   "Deploy project configuration file to remote host",
	Example: "forge deploy <project> -f <path/to/deployment.yml>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		secrets, err := vault.API.Get(ctx, args[0])
		if err != nil {
			errors.VaultErrors(err)
		}
		value, exists := secrets.Data["host"]
		if !exists {
			errors.VaultErrors(fmt.Errorf("host not found"))
		}
		host := remote.NewHost()
		host.Parse(value.(map[string]any))
		ssh, err := remote.NewSSH(vault.ENV, host.Addr)
		if err != nil {
			errors.RemoteErrors(err)
		}
		defer ssh.Close()
		localFile := filepath.Join(config.CONFIG_PATH, "config.yaml")
		remoteFile := filepath.Join(host.Path, "config.yaml")
		if err := ssh.Upload(localFile, remoteFile); err != nil {
			errors.RemoteErrors(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	generateCmd.Flags().StringVarP(&file, "file", "f", "", "path/to/deployment.yml")
	generateCmd.MarkFlagRequired("file")
}
