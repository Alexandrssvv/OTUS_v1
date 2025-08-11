package tag

type Tag struct {
	id   int    // Уникальный идентификатор тега.
	name string // Наименование тега, вводимое пользователем.
}

// New Конструктор, для создания пользователей (User)
func New(id int, name string) *Tag {
	return &Tag{
		id:   id,
		name: name,
	}
}

func (t *Tag) GetID() int {
	return t.id //Реализация интерфейса
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
