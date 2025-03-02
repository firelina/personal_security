package models

import "time"

type CreateReminderRequest struct {
	EventID            int       `json:"event_id"`
	ReminderTime       time.Time `json:"reminder_time"`
	NotificationMethod string    `json:"notification_method"`
}
