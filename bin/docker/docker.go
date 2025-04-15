// Package docker реализует функции для взаимодействия с Docker.
package docker

type Docker struct {
	client *docker.Client
}
