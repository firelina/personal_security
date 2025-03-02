package domain

import "time"

type Reminder struct {
	ID                 int       `json:"id"`
	EventID            int       `json:"event_id"`
	ReminderTime       time.Time `json:"reminder_time"`
	NotificationMethod string    `json:"notification_method"`
}
