package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_security/internal/domain"
	"personal_security/internal/gateways/http/models"
	"personal_security/internal/usecase"
)

type EventHandler struct {
	eventUsecase *usecase.Event
}

func NewEventHandler(u *usecase.Event) *EventHandler {
	return &EventHandler{u}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var newEvent domain.Event
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	eventID := h.eventUsecase.CreateEvent(&newEvent)
	c.JSON(http.StatusCreated, eventID)
}

func (h *EventHandler) GetEvents(c *gin.Context) {
	c.JSON(http.StatusOK, h.eventUsecase.GetEvents())
}

func (h *EventHandler) UpdateEventStatus(c *gin.Context) {
	var updateRequest models.UpdateRequest

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := h.eventUsecase.UpdateEventStatus(updateRequest)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Событие не найдено"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Статус события обновлен", "event": event})

}
