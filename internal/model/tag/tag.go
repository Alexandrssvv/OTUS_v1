package tag

type Tag struct {
	id   int    // Уникальный идентификатор тега.
	name string // Наименование тега, вводимое пользователем.
}

// Геттеры
func (t *Tag) ID() int {
	return t.id
}
func (t *Tag) Name() string {
	return t.name
}

// Сеттеры
func (t *Tag) SetName(name string) {
	t.name = name
}
