package user

// информация о зарегистрированных пользователях
import "time"

type User struct {
	id             int       //Уникальный ID пользователя
	email          string    //Адрес электронной почты
	hashedPassword string    //Хэш пароля (безопасное хранение пароля)
	createdAt      time.Time //Дата и время создания/регистрации пользователя
}

// Геттеры
func (u *User) ID() int {
	return u.id
}
func (u *User) Email() string {
	return u.email
}
func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

// Сеттеры
func (u *User) SetEmail(email string) {
	u.email = email
}
func (u *User) SetHashedPassword(hash string) {
	u.hashedPassword = hash
}

// Специальный метод (без экспонирования поля)
func (u *User) CheckPassword(hash string) bool {
	return u.hashedPassword == hash
}
