package remote

import (
	"testing"

	"github.com/PavelMilanov/forge/config"
)

func TestNewClient(t *testing.T) {
	env, err := config.NewEnv(config.CONFIG_PATH, config.CONFIG_FILE)
	if err != nil {
		t.Fatal(err)
	}
	c, err := NewClient(env)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
}
