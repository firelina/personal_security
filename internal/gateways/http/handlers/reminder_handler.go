package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_security/internal/domain"
	"personal_security/internal/gateways/http/models"
	"personal_security/internal/usecase"
	"strconv"
)

type ReminderHandler struct {
	reminderUsecase *usecase.ReminderService
}

func NewReminderHandler(u *usecase.ReminderService) *ReminderHandler {
	return &ReminderHandler{u}
}

func (h *ReminderHandler) CreateReminder(c *gin.Context) {
	var newReminder models.CreateReminderRequest
	if err := c.ShouldBindJSON(&newReminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	reminderID, err := h.reminderUsecase.CreateReminder(ctx, &domain.Reminder{
		EventID:            newReminder.EventID,
		ReminderTime:       newReminder.ReminderTime,
		NotificationMethod: newReminder.NotificationMethod,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusCreated, reminderID)
}

func (h *ReminderHandler) GetReminders(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("event_id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	reminders, err := h.reminderUsecase.GetReminders(ctx, eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, models.NewReminderResponse(reminders))
}

func (h *ReminderHandler) SendReminders(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	err = h.reminderUsecase.SendReminders(ctx, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Напоминания отправлены"})

}
