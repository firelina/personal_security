package usecase

import (
	"context"
	"personal_security/internal/domain"
	"personal_security/internal/repository"
)

type EventServiceInterface interface {
	CreateEvent(ctx context.Context, newEvent *domain.Event) (int, error)
	GetEvents(ctx context.Context, userID int) ([]*domain.Event, error)
	UpdateEventStatus(ctx context.Context, eventID int, status string) (*domain.Event, error)
}

type EventService struct {
	repo repository.EventRepositoryInterface
}

func NewEventService(repo repository.EventRepositoryInterface) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(ctx context.Context, newEvent *domain.Event) (int, error) {
	resultChan := make(chan struct {
		ID  int
		Err error
	})

	go func() {
		userID, err := s.repo.CreateEvent(ctx, newEvent)
		resultChan <- struct {
			ID  int
			Err error
		}{ID: userID, Err: err}
	}()

	result := <-resultChan
	return result.ID, result.Err
}

func (s *EventService) GetEvents(ctx context.Context, userID int) ([]*domain.Event, error) {
	resultChan := make(chan struct {
		Events []*domain.Event
		Err    error
	})

	go func() {
		events, err := s.repo.GetEvents(ctx, userID)
		resultChan <- struct {
			Events []*domain.Event
			Err    error
		}{Events: events, Err: err}
	}()

	result := <-resultChan
	return result.Events, result.Err
}

func (s *EventService) UpdateEventStatus(ctx context.Context, eventID int, status string) (*domain.Event, error) {
	resultChan := make(chan struct {
		Event *domain.Event
		Err   error
	})

	go func() {
		event, err := s.repo.UpdateEventStatus(ctx, eventID, status)
		resultChan <- struct {
			Event *domain.Event
			Err   error
		}{Event: event, Err: err}
	}()

	result := <-resultChan
	return result.Event, result.Err
}
