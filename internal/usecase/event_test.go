package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"personal_security/internal/domain"
	"testing"
	"time"
)

type MockEventRepository struct {
	mock.Mock
}

func (m *MockEventRepository) CreateEvent(ctx context.Context, newEvent *domain.Event) (int, error) {
	args := m.Called(ctx, newEvent)
	return args.Int(0), args.Error(1)
}

func (m *MockEventRepository) GetEvents(ctx context.Context, userID int) ([]*domain.Event, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*domain.Event), args.Error(1)
}

func (m *MockEventRepository) UpdateEventStatus(ctx context.Context, eventID int, status string) (*domain.Event, error) {
	args := m.Called(ctx, eventID, status)
	return args.Get(0).(*domain.Event), args.Error(1)
}

func TestCreateEvent(t *testing.T) {
	mockRepo := new(MockEventRepository)
	eventService := NewEventService(mockRepo)
	t.Run("success create event", func(t *testing.T) {
		newEvent := &domain.Event{UserID: 1, Title: "New Event", Date: time.Now(), Description: "This is a new event", Status: "pending"}

		mockRepo.On("CreateEvent", mock.Anything, newEvent).Return(1, nil)

		eventID, err := eventService.CreateEvent(context.Background(), newEvent)

		assert.NoError(t, err)
		assert.Equal(t, 1, eventID)

		mockRepo.AssertExpectations(t)
	})

}

func TestGetEvents(t *testing.T) {
	mockRepo := new(MockEventRepository)
	eventService := NewEventService(mockRepo)
	t.Run("success get events", func(t *testing.T) {
		userID := 1
		expectedEvents := []*domain.Event{
			{ID: 1, UserID: userID, Title: "Event 1", Date: time.Now(), Description: "This is event 1", Status: "pending"},
			{ID: 2, UserID: userID, Title: "Event 2", Date: time.Now(), Description: "This is event 2", Status: "pending"},
		}

		mockRepo.On("GetEvents", mock.Anything, userID).Return(expectedEvents, nil)

		events, err := eventService.GetEvents(context.Background(), userID)

		assert.NoError(t, err)
		assert.Equal(t, expectedEvents, events)

		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateEventStatus(t *testing.T) {
	mockRepo := new(MockEventRepository)
	eventService := NewEventService(mockRepo)
	t.Run("success update status", func(t *testing.T) {
		eventID := 1
		status := "sent"
		expectedEvent := &domain.Event{ID: eventID, UserID: 1, Title: "Event 1", Date: time.Now(), Description: "This is event 1", Status: status}

		mockRepo.On("UpdateEventStatus", mock.Anything, eventID, status).Return(expectedEvent, nil)

		event, err := eventService.UpdateEventStatus(context.Background(), eventID, status)

		assert.NoError(t, err)
		assert.Equal(t, expectedEvent, event)

		mockRepo.AssertExpectations(t)
	})
	t.Run("fail event update no found error", func(t *testing.T) {
		eventID := 1
		status := "approved"

		mockRepo.On("UpdateEventStatus", mock.Anything, eventID, status).Return(&domain.Event{}, errors.New("event not found"))

		event, err := eventService.UpdateEventStatus(context.Background(), eventID, status)

		assert.Error(t, err)
		assert.Equal(t, event, &domain.Event{})
		assert.Equal(t, "event not found", err.Error())

		mockRepo.AssertExpectations(t)
	})
}
