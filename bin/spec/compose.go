package spec

/*
Compose - абстракция над docker compose спецификацией в файлах конфигурации.
*/
type Compose struct {
	Tag string
}

func (c *Compose) Init() error {

	return nil
}
