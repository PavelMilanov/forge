package spec

type Kubernetes struct {
}

/*
Init инициализирует модель kubernetes значениями по умолчанию.
*/
func (k *Kubernetes) Init() {

}

/*
Generate генерирует конфигурацию kubernetes на основе шаблона.

Params

	path - имя шаблона.

Returns

	string - итоговое содержимое конфигурации.
	error - ошибка генерации.
*/
func (k *Kubernetes) Generate(path string) (string, error) {
	return "", nil
}

/*
Parse заполняет модель kubernetes данными из хранилища.
*/
func (k *Kubernetes) Parse(map[string]any) {
}

/*
Update применяет обновления параметров модели kubernetes.

Params

	data - входящие параметры в формате key=value.

Returns

	error - ошибка валидации или применения параметров.
*/
func (k *Kubernetes) Update(data []string) error {
	return nil
}

/*
Print выводит текущие значения модели kubernetes.
*/
func (k *Kubernetes) Print() {

}
