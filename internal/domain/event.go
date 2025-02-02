package domain

import "time"

type Event struct {
	ID          int        `json:"id"`          // Уникальный идентификатор события
	UserID      int        `json:"user_id"`     // Идентификатор пользователя, которому принадлежит событие
	Title       string     `json:"title"`       // Название события
	Date        time.Time  `json:"date"`        // Дата события
	Description string     `json:"description"` // Описание события
	Status      string     `json:"status"`      // Статус события (например, "запланировано", "завершено")
	Reminders   []Reminder `json:"reminders"`   // Список напоминаний, связанных с событием
}
