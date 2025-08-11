package service

import (
	"fmt"
	"notes-app/internal/model/note"
	"notes-app/internal/model/note_tag"
	"notes-app/internal/model/tag"
	"notes-app/internal/model/user"
	"notes-app/internal/repositoty"
	"time"
)

// GenerateData — создаёт User, Note, Tag и NoteTag и отправляет в репозиторий
func GenerateData() {
	now := time.Now()

	// Создание ID
	u := user.New(1, "user@example.com", "bcrypt:hash...", now)
	n := note.New(1, u.GetID(), "Первая заметка", "Текст заметки", now)
	t := tag.New(1, "Важно")

	// Связь между заметкой и тегом
	link := note_tag.New(1, n.GetID(), t.GetID())

	// Сохраняем в репозиторий
	repositoty.StoreBatch(u, n, t, link)

	fmt.Println("[service] данные успешно созданы и сохранены")
}
