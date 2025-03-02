package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_security/internal/domain"
	"personal_security/internal/gateways/http/models"
	"personal_security/internal/usecase"
	"strconv"
)

type EventHandler struct {
	eventUsecase *usecase.EventService
}

func NewEventHandler(u *usecase.EventService) *EventHandler {
	return &EventHandler{u}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var newEvent models.CreateEventRequest
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	eventID, err := h.eventUsecase.CreateEvent(ctx, &domain.Event{
		UserID:      newEvent.UserID,
		Title:       newEvent.Title,
		Date:        newEvent.Date,
		Description: newEvent.Description,
		Status:      newEvent.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusCreated, eventID)
}

func (h *EventHandler) GetEvents(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	events, err := h.eventUsecase.GetEvents(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, models.NewEventResponse(events))
}

func (h *EventHandler) UpdateEventStatus(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	var updateRequest models.UpdateRequest

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	event, err := h.eventUsecase.UpdateEventStatus(ctx, eventID, updateRequest.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Статус события обновлен", "event": models.NewEventResponseItem(event)})

}
