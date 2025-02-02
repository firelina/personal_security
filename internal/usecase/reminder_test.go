package usecase

import (
	"github.com/stretchr/testify/assert"
	"personal_security/internal/domain"
	"testing"
	"time"
)

func TestCreateReminder(t *testing.T) {
	reminderManager := &Reminder{}

	newReminder := &domain.Reminder{
		EventID:            1,
		ReminderTime:       time.Now().Add(1 * time.Hour), // Напоминание через 1 час
		NotificationMethod: "email",
	}

	reminderID := reminderManager.CreateReminder(newReminder)

	if reminderID != 1 {
		t.Errorf("Expected reminder ID to be 1, got %d", reminderID)
	}
	assert.Equal(t, 1, reminderID)
	assert.Equal(t, 1, len(reminders))
}

func TestGetReminders(t *testing.T) {
	reminderManager := &Reminder{}

	reminderManager.CreateReminder(&domain.Reminder{
		ReminderTime:       time.Now().Add(1 * time.Hour),
		NotificationMethod: "email",
	})

	remindersList := reminderManager.GetReminders()

	assert.Equal(t, 2, len(remindersList))
	assert.Equal(t, "email", remindersList[0].NotificationMethod)
}

func TestSendReminders(t *testing.T) {
	reminderManager := &Reminder{}

	event := &domain.Event{
		ID:    1,
		Title: "Встреча с клиентом",
		Reminders: []domain.Reminder{
			{
				EventID:            1,
				ReminderTime:       time.Now().Add(1 * time.Hour), // Напоминание через 1 час
				NotificationMethod: "email",
			},
		},
	}

	events = append(events, event)

	err := reminderManager.SendReminders()

	assert.Nil(t, err)

	events = nil
	err = reminderManager.SendReminders()

	assert.NotNil(t, err)
}
