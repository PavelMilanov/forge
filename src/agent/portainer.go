package agent

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/PavelMilanov/forge/api"
	"github.com/PavelMilanov/forge/config"
)

type PortainerAgent struct {
	portainer *api.Portainer
}

func NewPortainerAgent() *PortainerAgent {
	return &PortainerAgent{}
}

func (a *PortainerAgent) DeployStack(name string) error {
	portainer, err := api.NewPortainer(
		config.AppConfig.Agent.Credentials.Url,
		config.AppConfig.Agent.Credentials.Key)
	if err != nil {
		return err
	}
	stacks, err := portainer.GetStacks()
	if err != nil {
		return err
	}
	if len(stacks) == 0 {
		fmt.Println("no stacks found")
		return nil
	}
	for _, stack := range stacks {
		if stack.StackName == name {
			project, err := portainer.GetStackFile(stack)
			if err != nil {
				return err
			}
			if err := portainer.UpdateStack(project); err != nil {
				return err
			}
			fmt.Println("Stack updated")
			return nil
		}
		// fmt.Println("no stacks found")
	}
	return nil
}

func (a *PortainerAgent) CreateStack(stackName string, templateName string) error {
	portainer, err := api.NewPortainer(
		config.AppConfig.Agent.Credentials.Url,
		config.AppConfig.Agent.Credentials.Key)
	if err != nil {
		return err
	}
	tmpls, err := portainer.GetTemplates()
	if err != nil {
		return err
	}
	if len(tmpls) == 0 {
		fmt.Println("no stacks found")
		return nil
	}
	for _, tmpl := range tmpls {
		if tmpl.TemplateName == templateName {
			proj, err := portainer.GetTemplateFile(tmpl)
			if err != nil {
				return err
			}
			var buf bytes.Buffer
			data := map[string]string{
				"stand": stackName,
			}
			funcMap := template.FuncMap{
				"port": func(s string) string {
					return strings.TrimPrefix(s, "f") // Удаляет "f", если она есть в начале
				},
			}
			t := template.Must(template.New("portainer").Option("missingkey=error").Funcs(funcMap).Parse(string(proj.TemplateFile)))
			err = t.Execute(&buf, data)
			if err != nil {
				return err
			}
			if err := portainer.CreateStack(stackName+"-"+templateName, buf.String()); err != nil {
				return err
			}
			// fmt.Println("Stack created")
			return nil
		}
	}
	fmt.Println("no templates found")
	return nil
}
