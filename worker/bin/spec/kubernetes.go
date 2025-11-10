package spec

type Kubernetes struct {
}

func (k *Kubernetes) Init() {

}

func (k *Kubernetes) Generate(path string) (string, error) {
	return "", nil
}

func (k *Kubernetes) Parse(map[string]any) {
}

func (k *Kubernetes) Update(data []string) error {
	return nil
}

func (k *Kubernetes) Print() {

}
