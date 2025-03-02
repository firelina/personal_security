package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"personal_security/internal/domain"
	"testing"
	"time"
)

type MockReminderRepository struct {
	mock.Mock
}

func (m *MockReminderRepository) CreateReminder(ctx context.Context, newReminder *domain.Reminder) (int, error) {
	args := m.Called(ctx, newReminder)
	return args.Int(0), args.Error(1)
}

func (m *MockReminderRepository) GetReminders(ctx context.Context, eventID int) ([]*domain.Reminder, error) {
	args := m.Called(ctx, eventID)
	return args.Get(0).([]*domain.Reminder), args.Error(1)
}

type MockEventService struct {
	mock.Mock
}

func (m *MockEventService) GetEvents(ctx context.Context, userID int) ([]*domain.Event, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*domain.Event), args.Error(1)
}

func (m *MockEventService) CreateEvent(ctx context.Context, newEvent *domain.Event) (int, error) {
	return 0, nil
}

func (m *MockEventService) UpdateEventStatus(ctx context.Context, eventID int, status string) (*domain.Event, error) {
	return nil, nil
}

func TestCreateReminder(t *testing.T) {
	mockRepo := new(MockReminderRepository)
	reminderService := NewReminderService(mockRepo, nil)
	t.Run("success create reminder", func(t *testing.T) {
		newReminder := &domain.Reminder{EventID: 1, ReminderTime: time.Now().Add(1 * time.Hour), NotificationMethod: "email"}

		mockRepo.On("CreateReminder", mock.Anything, newReminder).Return(1, nil)

		reminderID, err := reminderService.CreateReminder(context.Background(), newReminder)

		assert.NoError(t, err)
		assert.Equal(t, 1, reminderID)

		mockRepo.AssertExpectations(t)
	})
}

func TestGetReminders(t *testing.T) {
	mockRepo := new(MockReminderRepository)
	reminderService := NewReminderService(mockRepo, nil)

	t.Run("success get reminders", func(t *testing.T) {
		eventID := 1
		expectedReminders := []*domain.Reminder{
			{ID: 1, EventID: eventID, ReminderTime: time.Now().Add(1 * time.Hour), NotificationMethod: "email"},
			{ID: 2, EventID: eventID, ReminderTime: time.Now().Add(2 * time.Hour), NotificationMethod: "sms"},
		}

		mockRepo.On("GetReminders", mock.Anything, eventID).Return(expectedReminders, nil)

		reminders, err := reminderService.GetReminders(context.Background(), eventID)

		assert.NoError(t, err)
		assert.Equal(t, expectedReminders, reminders)

		mockRepo.AssertExpectations(t)
	})
}

func TestSendReminders(t *testing.T) {
	mockRepo := new(MockReminderRepository)
	mockEventService := new(MockEventService)
	reminderService := NewReminderService(mockRepo, mockEventService)
	currentTime := time.Now()

	t.Run("success send reminder", func(t *testing.T) {
		userID := 1
		expectedEvents := []*domain.Event{
			{ID: 1, UserID: userID, Title: "Event 1", Date: currentTime.Add(2 * time.Hour), Description: "This is event 1", Status: "pending"},
		}

		mockEventService.On("GetEvents", mock.Anything, userID).Return(expectedEvents, nil)

		expectedReminders := []*domain.Reminder{
			{ID: 1, EventID: 1, ReminderTime: currentTime.Add(3 * time.Hour), NotificationMethod: "email"},
		}

		mockRepo.On("GetReminders", mock.Anything, 1).Return(expectedReminders, nil)

		err := reminderService.SendReminders(context.Background(), userID)

		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
		mockEventService.AssertExpectations(t)
	})

	t.Run("fail send reminder no reminder found", func(t *testing.T) {
		userID := 2
		expectedEvents := []*domain.Event{
			{ID: 2, UserID: userID, Title: "Event 1", Date: currentTime.Add(-2 * time.Hour), Description: "This is event 1", Status: "pending"},
		}

		mockEventService.On("GetEvents", mock.Anything, userID).Return(expectedEvents, nil)
		expectedReminders := []*domain.Reminder{
			{ID: 2, EventID: 2, ReminderTime: currentTime.Add(-1 * time.Hour), NotificationMethod: "email"},
		}

		mockRepo.On("GetReminders", mock.Anything, 2).Return(expectedReminders, nil)

		err := reminderService.SendReminders(context.Background(), userID)

		assert.Error(t, err)
		assert.Equal(t, "events happen before reminder time not found", err.Error())

		mockRepo.AssertExpectations(t)
		mockEventService.AssertExpectations(t)
	})
}
