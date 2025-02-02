package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_security/internal/domain"
	"personal_security/internal/usecase"
)

type ReminderHandler struct {
	reminderUsecase *usecase.Reminder
}

func NewReminderHandler(u *usecase.Reminder) *ReminderHandler {
	return &ReminderHandler{u}
}

func (h *ReminderHandler) CreateReminder(c *gin.Context) {
	var newReminder domain.Reminder
	if err := c.ShouldBindJSON(&newReminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reminderID := h.reminderUsecase.CreateReminder(&newReminder)
	c.JSON(http.StatusCreated, reminderID)
}

func (h *ReminderHandler) GetReminders(c *gin.Context) {
	c.JSON(http.StatusOK, h.reminderUsecase.GetReminders())
}

func (h *ReminderHandler) SendReminders(c *gin.Context) {
	err := h.reminderUsecase.SendReminders()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "События, для которых можно отправить напоминания не найдены"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Напоминания отправлены"})

}
