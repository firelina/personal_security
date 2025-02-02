package usecase

import (
	"github.com/stretchr/testify/assert"
	"personal_security/internal/domain"
	"personal_security/internal/gateways/http/models"
	"testing"
	"time"
)

func TestCreateEvent(t *testing.T) {
	eventManager := &Event{}

	newEvent := &domain.Event{
		Title:       "Встреча с клиентом",
		Date:        time.Now(),
		Description: "Обсуждение условий контракта",
		Status:      "запланировано",
	}

	eventID := eventManager.CreateEvent(newEvent)

	assert.Equal(t, 1, eventID)

	assert.Equal(t, 1, len(events))
}

func TestGetEvents(t *testing.T) {
	eventManager := &Event{}

	eventManager.CreateEvent(&domain.Event{
		Title:       "Встреча с клиентом",
		Date:        time.Now(),
		Description: "Обсуждение условий контракта",
		Status:      "запланировано",
	})

	eventsList := eventManager.GetEvents()

	assert.Equal(t, 2, len(eventsList))
	assert.Equal(t, "Встреча с клиентом", eventsList[0].Title)
}

func TestUpdateEventStatus(t *testing.T) {
	eventManager := &Event{}

	eventID := eventManager.CreateEvent(&domain.Event{
		Title:       "Встреча с клиентом",
		Date:        time.Now(),
		Description: "Обсуждение условий контракта",
		Status:      "запланировано",
	})

	updateRequest := models.UpdateRequest{
		EventID: eventID,
		Status:  "завершено",
	}

	updatedEvent, err := eventManager.UpdateEventStatus(updateRequest)

	assert.Nil(t, err)
	assert.Equal(t, "завершено", updatedEvent.Status)

	// Проверка на обновление несуществующего события
	updateRequest.EventID = 4
	_, err = eventManager.UpdateEventStatus(updateRequest)

	assert.NotNil(t, err)
}
