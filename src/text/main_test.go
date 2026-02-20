package text

import (
	"os"
	"testing"
	"text/template"
)

func TestMain(t *testing.T) {
	x := `This is {{ .Name }} and he is {{ .Age }} years old`
	ready_template, err := template.New("test").Parse(x)
	if err != nil {
		t.Errorf("Error parsing template: %v", err)
	}
	read := ready_template.ExecuteTemplate(os.Stdout, "test", struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  30,
	})
	if read != nil {
		t.Errorf("Error executing template: %v", err)
	}
}
