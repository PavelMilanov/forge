// Package docker реализует функции для взаимодействия с Docker.
package docker

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func DockerCommand(command, env, filepath string, a ...string) error {
	var cmd *exec.Cmd
	switch command {
	case "up":
		if len(a) == 0 {
			cmd = exec.Command("docker", "--context", env, "compose", "-f", filepath, "up", "-d")
		} else {
			cmd = exec.Command("docker", "--context", env, "compose", "-f", filepath, "up", "-d", strings.Join(a, " "))
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error %w", err)
		}
		return nil
	case "update":
		if len(a) == 0 {
			cmd = exec.Command("docker", "--context", env, "compose", "-f", filepath, "up", "-d", "--force-recreate")
		} else {
			cmd = exec.Command("docker", "--context", env, "compose", "-f", filepath, "up", "-d", strings.Join(a, " "), "--force-recreate")
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error %w", err)
		}
		return nil
	case "down":
		if len(a) == 0 {
			cmd = exec.Command("docker", "--context", env, "compose", "-f", filepath, "down")
		} else {
			cmd = exec.Command("docker", "--context", env, "compose", "-f", filepath, "down", strings.Join(a, " "))
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error %w", err)
		}
		return nil
	default:
		return errors.New("unknown command")
	}
}
