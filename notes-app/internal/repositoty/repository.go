package repositoty

import (
	"fmt"
	"notes-app/internal/model/note"
	"notes-app/internal/model/note_tag"
	"notes-app/internal/model/tag"
	"notes-app/internal/model/user"
)

// NotesAppIdentifiable — идентифицируемая сущность приложения "Заметки". Объявление интерфейса.
type NotesAppIdentifiable interface {
	GetID() int
}

// "Хранилище в памяти" - различные слайсы
var (
	Users    = make([]user.User, 0, 16)
	Notes    = make([]note.Note, 0, 64)
	Tags     = make([]tag.Tag, 0, 32)
	NoteTags = make([]note_tag.NoteTag, 0, 128)

	typeCounters = map[string]int{
		"user":     0,
		"note":     0,
		"tag":      0,
		"note_tag": 0,
	}
)

// Store - раскладываем по слайсам
func Store(e NotesAppIdentifiable) error {
	switch v := e.(type) {
	case *user.User:
		Users = append(Users, *v)
		typeCounters["user"]++
		fmt.Printf("[repo] stored User id=%d (total users=%d)\n", v.GetID(), typeCounters["user"])
		return nil
	case *note.Note:
		Notes = append(Notes, *v)
		typeCounters["note"]++
		fmt.Printf("[repo] stored Note id=%d (total notes=%d)\n", v.GetID(), typeCounters["note"])
		return nil
	case *tag.Tag:
		Tags = append(Tags, *v)
		typeCounters["tag"]++
		fmt.Printf("[repo] stored Tag id=%d (total tags=%d)\n", v.GetID(), typeCounters["tag"])
		return nil
	case *note_tag.NoteTag:
		NoteTags = append(NoteTags, *v)
		typeCounters["note_tag"]++
		fmt.Printf("[repo] stored NoteTag id=%d (note=%d, tag=%d, total links=%d)\n", v.GetID(), v.NoteID(), v.TagID(), typeCounters["note_tag"])
		return nil
	default:
		return fmt.Errorf("Store: unsupported type %T", e)
	}
}

// StoreBatch — пакетная запись разных структур через интерфейс.
func StoreBatch(items ...NotesAppIdentifiable) error {
	for _, it := range items {
		if err := Store(it); err != nil {
			return fmt.Errorf("StoreBatch: failed to store item: %w", err)
		}
	}
	return nil
}

// Counters — возвращает текущие значения счётчиков.
func Counters() map[string]int {
	return typeCounters
}
