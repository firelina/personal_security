package models

import (
	"personal_security/internal/domain"
	"time"
)

type ReminderResponseItem struct {
	ID                 int       `json:"id"`
	EventID            int       `json:"event_id"`
	ReminderTime       time.Time `json:"reminder_time"`
	NotificationMethod string    `json:"notification_method"`
}

type ReminderResponse []ReminderResponseItem

func NewReminderResponse(reminders []*domain.Reminder) ReminderResponse {
	var response []ReminderResponseItem
	for _, r := range reminders {
		response = append(response, ReminderResponseItem{
			ID:                 r.ID,
			EventID:            r.EventID,
			ReminderTime:       r.ReminderTime,
			NotificationMethod: r.NotificationMethod,
		})
	}
	return response
}
