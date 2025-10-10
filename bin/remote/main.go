package remote

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/PavelMilanov/forge/config"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

type SSH struct {
	Client *ssh.Client
	File   *sftp.Client
}

func NewSSH(env *config.Env, server string) (*SSH, error) {
	var idx int
	for i, alias := range env.SSH.Servers {
		if alias.Host == server {
			idx = i
			break
		}
	}
	knowHostsPath := filepath.Dir(env.SSH.PrivateKey)
	knowHostsFile := filepath.Join(knowHostsPath, "known_hosts")
	key, err := os.ReadFile(env.SSH.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key: %v", err)
	}
	hostKeyCallback, err := knownhosts.New(knowHostsFile)
	if err != nil {
		return nil, fmt.Errorf("unable to load known_hosts: %w", err)
	}
	config := &ssh.ClientConfig{
		User: env.SSH.Servers[idx].User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallback,
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", env.SSH.Servers[idx].Server, 22), config)
	if err != nil {
		sshClient.Close()
		return nil, fmt.Errorf("unable to connect: %v", err)
	}
	fileClient, err := sftp.NewClient(sshClient)
	if err != nil {
		fileClient.Close()
		return nil, fmt.Errorf("unable to create sftp client: %w", err)
	}
	return &SSH{sshClient, fileClient}, nil
}
