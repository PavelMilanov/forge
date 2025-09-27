package remote

import (
	"bytes"
	"fmt"
	"io"
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

func NewSSH(env *config.Env, addr string) (*SSH, error) {
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
		User: env.SSH.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallback,
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", addr, 22), config)
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

func (ssh *SSH) Close() error {
	if ssh.File != nil {
		ssh.File.Close()
	}
	if ssh.Client != nil {
		return ssh.Client.Close()
	}
	return nil
}

func (ssh *SSH) RunCommand(cmd string) (string, string, error) {
	if ssh.Client == nil {
		return "", "", fmt.Errorf("ssh client is not initialized")
	}

	// создаём новую сессию
	session, err := ssh.Client.NewSession()
	if err != nil {
		return "", "", fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	// выполняем команду
	if err := session.Run(cmd); err != nil {
		return stdout.String(), stderr.String(), fmt.Errorf("failed to run command: %w", err)
	}

	return stdout.String(), stderr.String(), nil
}

// Upload копирует локальный файл на сервер
func (ssh *SSH) Upload(localPath, remotePath string) error {
	if ssh.File == nil {
		return fmt.Errorf("sftp client not initialized")
	}

	srcFile, err := os.Open(localPath)
	if err != nil {
		return fmt.Errorf("unable to open local file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := ssh.File.Create(remotePath)
	if err != nil {
		return fmt.Errorf("unable to create remote file: %w", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// Download копирует файл с сервера на локальную машину
func (ssh *SSH) Download(remotePath, localPath string) error {
	if ssh.File == nil {
		return fmt.Errorf("sftp client not initialized")
	}

	srcFile, err := ssh.File.Open(remotePath)
	if err != nil {
		return fmt.Errorf("unable to open remote file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(localPath)
	if err != nil {
		return fmt.Errorf("unable to create local file: %w", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}
