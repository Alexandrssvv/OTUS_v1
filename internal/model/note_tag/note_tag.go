package notetag

type NoteTag struct {
	noteID int // ID заметки (note)
	tagID  int // ID тега (tag)
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
