package usecase

import (
	"errors"
	"log"
	"personal_security/internal/domain"
	"time"
)

type Reminder struct {
}

var reminders []*domain.Reminder

func (r *Reminder) CreateReminder(newReminder *domain.Reminder) int {
	newReminder.ID = len(contacts) + 1
	reminders = append(reminders, newReminder)
	return newReminder.ID
}

func (r *Reminder) GetReminders() []*domain.Reminder {
	return reminders
}

func (r *Reminder) SendReminders() error {
	currentTime := time.Now()

	var upcomingReminders []domain.Reminder

	for _, event := range events {
		for _, reminder := range event.Reminders {
			if reminder.ReminderTime.After(currentTime) {
				upcomingReminders = append(upcomingReminders, reminder)
			}
		}
	}

	if len(upcomingReminders) == 0 {
		return errors.New("events happen before reminder time not found")
	}

	for _, reminder := range upcomingReminders {
		log.Printf("Отправка %s напоминания для события: %d, запланированного на %s в %s",
			reminder.NotificationMethod,
			reminder.EventID,
			reminder.ReminderTime.Format("2006-01-02"),
			reminder.ReminderTime.Format("15:04:05"),
		)
	}
	return nil
}
