package note

import (
	"time"
)

type Note struct {
	id        int        //	Уникальный идентификатор заметки.
	userID    int        // ID пользователя, которому принадлежит заметка.
	title     string     // Заголовок заметки.
	content   string     // Основной текст заметки.
	expiresAt *time.Time // Дата и время, после которых заметка считается устаревшей (необязательное поле).
	createdAt time.Time  // Дата и время создания заметки.
	updatedAt time.Time  // Дата и время последнего обновления заметки.
}

// Геттеры
func (n *Note) ID() int {
	return n.id
}
func (n *Note) UserID() int {
	return n.userID
}
func (n *Note) Title() string {
	return n.title
}
func (n *Note) Content() string {
	return n.content
}
func (n *Note) CreatedAt() time.Time {
	return n.createdAt
}
func (n *Note) UpdatedAt() time.Time {
	return n.updatedAt
}
func (n *Note) ExpiresAt() *time.Time {
	return n.expiresAt
}

// Сеттеры
func (n *Note) SetUserID(uid int) {
	n.userID = uid
}
func (n *Note) SetContent(title, content string) {
	n.title = title
	n.content = content
}
func (n *Note) SetExpiration(exp *time.Time) {
	n.expiresAt = exp
}
