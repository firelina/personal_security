package models

import (
	"personal_security/internal/domain"
	"time"
)

type EventResponseItem struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}

func NewEventResponseItem(event *domain.Event) EventResponseItem {
	return EventResponseItem{
		ID:          event.ID,
		UserID:      event.UserID,
		Title:       event.Title,
		Date:        event.Date,
		Description: event.Description,
		Status:      event.Status,
	}
}

type EventResponse []EventResponseItem

func NewEventResponse(events []*domain.Event) EventResponse {
	var response []EventResponseItem
	for _, e := range events {
		response = append(response, NewEventResponseItem(e))
	}
	return response
}
