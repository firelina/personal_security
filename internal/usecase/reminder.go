package usecase

import (
	"context"
	"fmt"
	"log"
	"personal_security/internal/domain"
	"personal_security/internal/repository"
	"time"
)

type ReminderService struct {
	repo         repository.ReminderRepositoryInterface
	eventService EventServiceInterface
}

func NewReminderService(repo repository.ReminderRepositoryInterface, eventService EventServiceInterface) *ReminderService {
	return &ReminderService{
		repo:         repo,
		eventService: eventService,
	}
}

func (s *ReminderService) CreateReminder(ctx context.Context, newReminder *domain.Reminder) (int, error) {
	resultChan := make(chan struct {
		ID  int
		Err error
	})

	go func() {
		userID, err := s.repo.CreateReminder(ctx, newReminder)
		resultChan <- struct {
			ID  int
			Err error
		}{ID: userID, Err: err}
	}()

	result := <-resultChan
	return result.ID, result.Err
}

func (s *ReminderService) GetReminders(ctx context.Context, eventID int) ([]*domain.Reminder, error) {
	resultChan := make(chan struct {
		Reminders []*domain.Reminder
		Err       error
	})

	go func() {
		reminders, err := s.repo.GetReminders(ctx, eventID)
		resultChan <- struct {
			Reminders []*domain.Reminder
			Err       error
		}{Reminders: reminders, Err: err}
	}()

	result := <-resultChan
	return result.Reminders, result.Err
}

func (s *ReminderService) SendReminders(ctx context.Context, userID int) error {
	events, err := s.eventService.GetEvents(ctx, userID)
	if err != nil {
		return err
	}

	currentTime := time.Now()

	var upcomingReminders []*domain.Reminder

	for _, event := range events {
		reminders, err := s.GetReminders(ctx, event.ID)
		if err != nil {
			return err
		}
		for _, reminder := range reminders {
			if reminder.ReminderTime.After(currentTime) {
				upcomingReminders = append(upcomingReminders, reminder)
			}
		}
	}

	if len(upcomingReminders) == 0 {
		return fmt.Errorf("events happen before reminder time not found")
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
