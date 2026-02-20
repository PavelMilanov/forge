package models

type PortainerEndpoint struct {
	ID   int    `json:"Id"`   // id
	Name string `json:"Name"` // name
	Type int    `json:"Type"` // 1,2,3 - docker, swarm, kubernetes
	URL  string `json:"URL"`
}
