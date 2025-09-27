package remote

import (
	"testing"

	"github.com/PavelMilanov/forge/config"
)

func TestNewClient(t *testing.T) {
	addr := "localhost"
	user := "vagrant"
	env, err := config.NewEnv("../var/config", "forge.yml")
	if err != nil {
		t.Fatal(err)
	}
	c, err := NewClient(env, user, addr)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
}
