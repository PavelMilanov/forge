package templatevars

import (
	"testing"
)

func TestExtractFromYAML_Success(t *testing.T) {
	yml := []byte(`
services:
  api:
    image: "{{.api.image}}"
    deploy:
      replicas: "{{.api.replicas}}"
  worker:
    image: "{{.worker.image}}"
  beat:
    image: "{{.beat.image}}"
  openim:
    image: "{{.openim.image}}"
`)

	result, err := ExtractFromYAML(yml)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := map[string]any{
		"api.image":    "",
		"api.replicas": 1,
		"worker.image": "",
		"beat.image":   "",
		"openim.image": "",
	}

	if len(result.Variables) != len(expected) {
		t.Fatalf("unexpected variables count: got=%d want=%d", len(result.Variables), len(expected))
	}

	for k, want := range expected {
		got, ok := result.Variables[k]
		if !ok {
			t.Fatalf("missing key %q", k)
		}
		if got != want {
			t.Fatalf("unexpected value for %q: got=%v want=%v", k, got, want)
		}
	}
}

func TestExtractFromYAML_DeduplicatesAndMultipleMatches(t *testing.T) {
	yml := []byte(`
msg: "{{.api.image}} -> {{ .api.image }}"
services:
  api:
    image: "{{.api.image}}"
`)

	result, err := ExtractFromYAML(yml)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Variables) != 1 {
		t.Fatalf("unexpected variables count: got=%d want=1", len(result.Variables))
	}

	if result.Variables["api.image"] != "" {
		t.Fatalf("unexpected default for api.image: got=%v want=\"\"", result.Variables["api.image"])
	}
}

func TestExtractFromYAML_InvalidYAML(t *testing.T) {
	_, err := ExtractFromYAML([]byte("services:\n  api: ["))
	if err == nil {
		t.Fatal("expected error for invalid yaml, got nil")
	}
}

func TestExtractFromYAML_NoTemplateVars(t *testing.T) {
	yml := []byte(`
services:
  api:
    image: "registry.example.com/app:latest"
`)

	result, err := ExtractFromYAML(yml)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Variables) != 0 {
		t.Fatalf("unexpected variables count: got=%d want=0", len(result.Variables))
	}
}

func TestExtractFromYAML_WithControlDirectives(t *testing.T) {
	yml := []byte(`
services:
  api:
    deploy:
      {{- if .placement.enabled }}
      placement:
        constraints:
          - "{{.placement.constraint}}"
      {{- end }}
    image: "{{.api.image}}"
`)

	result, err := ExtractFromYAML(yml)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Variables["api.image"] != "" {
		t.Fatalf("unexpected api.image default: %v", result.Variables["api.image"])
	}
	if result.Variables["placement.constraint"] != "" {
		t.Fatalf("unexpected placement.constraint default: %v", result.Variables["placement.constraint"])
	}
}
