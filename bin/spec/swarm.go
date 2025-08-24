package spec

/*
Swarm - абстракция над docker swarm спецификацией в файлах конфигурации.
*/
type Swarm struct {
	Tag      string
	Replicas int
}

func (s *Swarm) Init() error {
	return nil
}
