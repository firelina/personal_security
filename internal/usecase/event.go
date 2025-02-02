package usecase

import (
	"errors"
	"personal_security/internal/domain"
	"personal_security/internal/gateways/http/models"
)

type Event struct {
}

var events []*domain.Event

func (e *Event) CreateEvent(newEvent *domain.Event) int {
	newEvent.ID = len(events) + 1
	events = append(events, newEvent)
	return newEvent.ID
}

func (e *Event) GetEvents() []*domain.Event {
	return events
}

func (e *Event) UpdateEventStatus(updateRequest models.UpdateRequest) (*domain.Event, error) {
	for i, event := range events {
		if event.ID == updateRequest.EventID {
			events[i].Status = updateRequest.Status

			return events[i], nil
		}
	}
	return nil, errors.New("event by id not found")
}
