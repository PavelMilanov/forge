package spec

import (
	"bytes"
	"html/template"
	"testing"
)

const swarmTmp1 = `
services:
  alpine:
    image: alpine
    tag: {{.Tag}}
    deployment:
      replicas: {{.Replicas}}

  nginx:
    image: nginx
    tag: {{.Tag}}
    deployment:
      replicas: {{.Replicas}}
`

const swarmTmp2 = `
services:
  alpine:
    image: alpine
    tag: {{.Tag}}

  nginx:
    image: nginx
    tag: {{.Tag}}
    deployment:
      replicas: {{.Replicas}}
`

func TestSwarm(t *testing.T) {
	swarm := Swarm{Tag: "latest", Replicas: 1}

	for _, tmpl := range []string{swarmTmp1, swarmTmp2} {
		tmpl, err := template.New("swarm-template").Parse(tmpl)
		if err != nil {
			t.Errorf("Ошибка парсинга шаблона: %v", err)
		}
		var buf bytes.Buffer

		err = tmpl.Execute(&buf, swarm)
		if err != nil {
			t.Error(err)
		}

		t.Log(buf.String())
	}

}
