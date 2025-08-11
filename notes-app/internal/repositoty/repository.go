package repositoty

import (
	"fmt"
	"notes-app/internal/model"
	"notes-app/internal/model/note"
	"notes-app/internal/model/note_tag"
	"notes-app/internal/model/tag"
	"notes-app/internal/model/user"
)

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
func Store(m model.Model) {
	switch v := m.(type) {
	case *user.User:
		Users = append(Users, *v)
		typeCounters["user"]++
		fmt.Printf("[repo] stored User id=%d (total users=%d)\n", v.GetID(), typeCounters["user"])
	case *note.Note:
		Notes = append(Notes, *v)
		typeCounters["note"]++
		fmt.Printf("[repo] stored Note id=%d (total notes=%d)\n", v.GetID(), typeCounters["note"])
	case *tag.Tag:
		Tags = append(Tags, *v)
		typeCounters["tag"]++
		fmt.Printf("[repo] stored Tag id=%d (total tags=%d)\n", v.GetID(), typeCounters["tag"])
	case *note_tag.NoteTag:
		NoteTags = append(NoteTags, *v)
		typeCounters["note_tag"]++
		fmt.Printf("[repo] stored NoteTag id=%d (note=%d, tag=%d, total links=%d)\n", v.GetID(), v.NoteID(), v.TagID(), typeCounters["note_tag"])
	default:
		fmt.Println("[repo] unknown model type — ignored")
	}
}

// StoreBatch — пакетная запись разных структур через интерфейс.
func StoreBatch(models ...model.Model) {
	for _, m := range models {
		Store(m)
	}
}

// Counters — возвращает текущие значения счётчиков.
func Counters() map[string]int {
	return typeCounters
}
