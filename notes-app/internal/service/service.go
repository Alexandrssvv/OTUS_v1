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
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	idCounter := 1

	for {
		select {
		case <-ticker.C:
			now := time.Now()

			// Создание ID
			u := user.New(idCounter, fmt.Sprintf("user%d@example.com", idCounter), "bcrypt:hash...", now)
			n := note.New(idCounter, u.GetID(), fmt.Sprintf("Заметка %d", idCounter), "Текст заметки", now)
			t := tag.New(idCounter, fmt.Sprintf("Тег%d", (idCounter%3)+1))

			// Связь между заметкой и тегом
			link := note_tag.New(idCounter, n.GetID(), t.GetID())

			// Сохраняем в репозиторий
			if err := repositoty.StoreBatch(u, n, t, link); err != nil {
				fmt.Println("store error:", err)
				return
			}

			fmt.Println("Данные успешно сгенерированы и сохранены с ID=%d\n", idCounter)
			fmt.Println("Counters:", repositoty.Counters())

			idCounter++
		}
	}
}
