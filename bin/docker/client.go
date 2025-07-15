package docker

import (
	"fmt"
	"os"
	"os/exec"
)

func RegistryLogin(env string, login, password, registry string) error {
	cmd := exec.Command("echo", password, "|", "docker", "--context", env, "login", "-u", login, "--password-stdin", registry)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error %w", err)
	}
	return nil
}
