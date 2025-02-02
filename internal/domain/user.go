package domain

type User struct {
	ID       int       `json:"id"`       // Уникальный идентификатор пользователя
	Name     string    `json:"name"`     // Имя пользователя
	Email    string    `json:"email"`    // Email пользователя
	Password string    `json:"password"` // Пароль пользователя (в хэшированном виде)
	Events   []Event   `json:"events"`   // Список событий, связанных с пользователем
	Contacts []Contact `json:"contacts"` // Список контактов, связанных с пользователем
}
