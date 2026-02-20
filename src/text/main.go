package text

import (
	"html/template"
	"os"

	"github.com/PavelMilanov/forge/models"
)

func PrintEndpoints(items []models.PortainerEndpoint) error {
	tmpls := `Endpoints:
{{ range . }}
   ID: {{ .ID }}
   Name: {{ .Name }}
   URL: {{ .URL }}
{{ end }}
`
	parsed_template, err := template.New("print").Parse(tmpls)
	if err != nil {
		return err
	}
	read := parsed_template.Execute(os.Stdout, items)
	if read != nil {
		return err
	}
	return nil
}
