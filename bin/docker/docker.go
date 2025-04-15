// Package docker реализует функции для взаимодействия с Docker.
package docker

type Docker struct {
}

// func (d *Docker) command(command string, a ...string) error {
// 	var cmd *exec.Cmd
// 	switch command {
// 	case "up":
// 		if len(a) == 0 {
// 			cmd = exec.Command("docker", "compose", "-f", c.File, "up", "-d")
// 		} else {
// 			cmd = exec.Command("docker", "compose", "-f", c.File, "up", "-d", strings.Join(a, " "))
// 		}
// 		cmd.Stdout = os.Stdout
// 		cmd.Stderr = os.Stderr
// 		if err := cmd.Run(); err != nil {
// 			return fmt.Errorf("error %w", err)
// 		}
// 		return nil
// 	case "down":
// 		if len(a) == 0 {
// 			cmd = exec.Command("docker", "compose", "-f", c.File, "down")
// 		} else {
// 			cmd = exec.Command("docker", "compose", "-f", c.File, "down", strings.Join(a, " "))
// 		}
// 		cmd.Stdout = os.Stdout
// 		cmd.Stderr = os.Stderr
// 		if err := cmd.Run(); err != nil {
// 			return fmt.Errorf("error %w", err)
// 		}
// 		return nil
// 	default:
// 		return errors.New("неизвестная команда")
// 	}
// }
