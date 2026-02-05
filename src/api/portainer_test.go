package api

import (
	"testing"

	"github.com/PavelMilanov/forge/config"
)

func preRun() *Portainer {
	env, err := config.NewEnv("../var/forge", "forge")
	if err != nil {
		return nil
	}
	portainer, err := NewPortainer(env.Agent.Credentials.Url, env.Agent.Credentials.Key)
	if err != nil {
		return nil
	}
	return portainer
}

func TestGetStacks(t *testing.T) {
	portainer := preRun()
	data, err := portainer.GetStacks()
	if err != nil {
		t.Errorf("Error getting stacks: %s", err)
	}
	t.Logf("%+v", data)
}

func TestGetTemplates(t *testing.T) {
	portainer := preRun()
	data, err := portainer.GetTemplates()
	if err != nil {
		t.Errorf("Error getting templates: %s", err)
	}
	t.Logf("%+v", data)
}

func TestGetStackFile(t *testing.T) {
	portainer := preRun()
	stacks, err := portainer.GetStacks()
	if err != nil {
		t.Errorf("Error getting stacks: %s", err)
	}
	data, err := portainer.GetStackFile(stacks[0])
	if err != nil {
		t.Errorf("Error getting stack file: %s", err)
	}
	t.Logf("%+v", data)
}

func TestGetTemplateFile(t *testing.T) {
	portainer := preRun()
	templates, err := portainer.GetTemplates()
	if err != nil {
		t.Errorf("Error getting templates: %s", err)
	}
	data, err := portainer.GetTemplateFile(templates[0])
	if err != nil {
		t.Errorf("Error getting template file: %s", err)
	}
	t.Logf("%+v", data)
}

func TestUpdateStack(t *testing.T) {
	portainer := preRun()
	stacks, err := portainer.GetStacks()
	if err != nil {
		t.Errorf("Error getting stacks: %s", err)
	}
	data, err := portainer.GetStackFile(stacks[0])
	if err != nil {
		t.Errorf("Error getting stack file: %s", err)
	}
	if err := portainer.UpdateStack(data); err != nil {
		t.Errorf("Error updating stack: %s", err)
	}
}
