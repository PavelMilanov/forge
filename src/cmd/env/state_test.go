package env

import "testing"

func TestSetPathIfMissing(t *testing.T) {
	data := map[string]any{
		"api": map[string]any{
			"image": "custom-image",
		},
	}

	setPathIfMissing(data, "api.image", "")
	setPathIfMissing(data, "api.replicas", 1)
	setPathIfMissing(data, "worker.image", "")

	api := data["api"].(map[string]any)
	if api["image"] != "custom-image" {
		t.Fatalf("existing value was overwritten: got=%v", api["image"])
	}
	if api["replicas"] != 1 {
		t.Fatalf("missing value was not set: got=%v", api["replicas"])
	}

	worker := data["worker"].(map[string]any)
	if worker["image"] != "" {
		t.Fatalf("worker.image unexpected value: got=%v", worker["image"])
	}
}

func TestEnsureEnvironmentPlacementDefaultFalse(t *testing.T) {
	state := projectState{
		Environments: map[string]environmentRef{},
	}

	ensureEnvironment(&state, "stage")
	ref := state.Environments["stage"]

	if ref.Placement["enabled"] != false {
		t.Fatalf("placement.enabled should be false by default, got=%v", ref.Placement["enabled"])
	}
}
