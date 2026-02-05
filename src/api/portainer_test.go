package api

import (
	"testing"

	"github.com/PavelMilanov/forge/config"
)

func TestGetStacks(t *testing.T) {
	env, err := config.NewEnv("../var/forge", "forge")
	if err != nil {
		t.Errorf("Error loading environment: %s", err)
	}
	portainer, err := NewPortainer(env.Agent.Credentials.Url, env.Agent.Credentials.Key)
	if err != nil {
		t.Errorf("Error creating Portainer client: %s", err)
	}
	data, err := portainer.GetStacks()
	if err != nil {
		t.Errorf("Error getting stacks: %s", err)
	}
	t.Logf("%+v", data)
}

func TestGetTemplates(t *testing.T) {
	env, err := config.NewEnv("../var/forge", "forge")
	if err != nil {
		t.Errorf("Error loading environment: %s", err)
	}
	portainer, err := NewPortainer(env.Agent.Credentials.Url, env.Agent.Credentials.Key)
	if err != nil {
		t.Errorf("Error creating Portainer client: %s", err)
	}
	data, err := portainer.GetTemplates()
	if err != nil {
		t.Errorf("Error getting templates: %s", err)
	}
	t.Logf("%+v", data)
}

func TestGetStackFile(t *testing.T) {
	env, err := config.NewEnv("../var/forge", "forge")
	if err != nil {
		t.Errorf("Error loading environment: %s", err)
	}
	portainer, err := NewPortainer(env.Agent.Credentials.Url, env.Agent.Credentials.Key)
	if err != nil {
		t.Errorf("Error creating Portainer client: %s", err)
	}
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
	env, err := config.NewEnv("../var/forge", "forge")
	if err != nil {
		t.Errorf("Error loading environment: %s", err)
	}
	portainer, err := NewPortainer(env.Agent.Credentials.Url, env.Agent.Credentials.Key)
	if err != nil {
		t.Errorf("Error creating Portainer client: %s", err)
	}
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
	env, err := config.NewEnv("../var/forge", "forge")
	if err != nil {
		t.Errorf("Error loading environment: %s", err)
	}
	portainer, err := NewPortainer(env.Agent.Credentials.Url, env.Agent.Credentials.Key)
	if err != nil {
		t.Errorf("Error creating Portainer client: %s", err)
	}
	stacks, err := portainer.GetStacks()
	if err != nil {
		t.Errorf("Error getting stacks: %s", err)
	}
	data, err := portainer.GetStackFile(stacks[0])
	if err != nil {
		t.Errorf("Error getting stack file: %s", err)
	}
	resp, err := portainer.UpdateStack(data)
	if err != nil {
		t.Errorf("Error updating stack: %s", err)
	}
	t.Logf("%s", resp)
}
