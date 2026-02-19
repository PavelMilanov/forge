package agent

type SSHAgent struct {
}

func NewSSHAgent() *SSHAgent {
	return &SSHAgent{}
}

func (a *SSHAgent) DeployStack(name string) error {
	// Implement stack deployment logic here
	return nil
}

func (a *SSHAgent) CreateStack(endpointId int, stackName string, template string) error {
	// Implement stack creation logic here
	return nil
}

func (a *SSHAgent) ListEndpoints() (map[int]string, error) {
	// Implement endpoint listing logic here
	return nil, nil
}
