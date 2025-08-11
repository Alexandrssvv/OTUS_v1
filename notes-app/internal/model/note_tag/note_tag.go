package note_tag

type NoteTag struct {
	id     int // Уникальный идентификатор записи
	noteID int // ID заметки (note)
	tagID  int // ID тега (tag)
}

// New Конструктор, для создания пользователей (User)
func New(id, noteID, tagID int) *NoteTag {
	return &NoteTag{
		id:     id,
		noteID: noteID,
		tagID:  tagID,
	}
}

func (nt *NoteTag) GetID() int {
	return nt.id //Реализация интерфейса
}

// Геттеры
func (nt *NoteTag) NoteID() int {
	return nt.noteID
}
func (nt *NoteTag) TagID() int {
	return nt.tagID
}

// Сеттеры
func (nt *NoteTag) Set(noteID, tagID int) {
	nt.noteID = noteID
	nt.tagID = tagID
}
