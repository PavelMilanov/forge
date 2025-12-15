package config

import (
	"testing"
)

func TestNewEnv(t *testing.T) {
	env, err := NewEnv("../var/config/", "forge.yml")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", env)
}
