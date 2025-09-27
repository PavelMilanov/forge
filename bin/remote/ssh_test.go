package remote

import (
	"testing"

	"github.com/PavelMilanov/forge/config"
)

func TestNewClient(t *testing.T) {
	addr := "localhost"
	env, err := config.NewEnv("../var/config", "forge.yml")
	if err != nil {
		t.Fatal(err)
	}
	c, err := NewSSH(env, addr)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
}
