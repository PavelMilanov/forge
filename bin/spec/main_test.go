package spec

import "testing"

func TestNewSpec(t *testing.T) {
	t.Run("swarm", func(t *testing.T) {
		_, err := NewSpec("swarm")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
	t.Run("compose", func(t *testing.T) {
		_, err := NewSpec("compose")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
	t.Run("kubernetes", func(t *testing.T) {
		_, err := NewSpec("kubernetes")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
	t.Run("docker", func(t *testing.T) {
		_, err := NewSpec("docker")
		if err.Error() != "unknown mode" {
			t.Errorf("unexpected error: %v", err)
		}

	})
	t.Run("minikube", func(t *testing.T) {
		_, err := NewSpec("minikube")
		if err.Error() != "unknown mode" {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
