package domain

type Contact struct {
	ID     int    `json:"id"`      // Уникальный идентификатор контакта
	UserID int    `json:"user_id"` // Идентификатор пользователя, которому принадлежит контакт
	Name   string `json:"name"`    // Имя контакта
	Phone  string `json:"phone"`   // Номер телефона контакта
	Email  string `json:"email"`   // Email контакта
}
