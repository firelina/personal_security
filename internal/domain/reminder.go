package domain

import "time"

type Reminder struct {
	ID                 int       `json:"id"`                  // Уникальный идентификатор напоминания
	EventID            int       `json:"event_id"`            // Идентификатор события, к которому относится напоминание
	ReminderTime       time.Time `json:"reminder_time"`       // Время, когда должно произойти напоминание
	NotificationMethod string    `json:"notification_method"` // Способ уведомления (например, "email", "push-уведомление")
}
