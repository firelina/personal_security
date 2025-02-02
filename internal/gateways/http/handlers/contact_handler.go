package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_security/internal/domain"
	"personal_security/internal/usecase"
)

type ContactHandler struct {
	contactUsecase *usecase.Contact
}

func NewContactHandler(u *usecase.Contact) *ContactHandler {
	return &ContactHandler{u}
}

func (h *ContactHandler) CreateEvent(c *gin.Context) {
	var newContact domain.Contact
	if err := c.ShouldBindJSON(&newContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contactID := h.contactUsecase.CreateContact(&newContact)
	c.JSON(http.StatusCreated, contactID)
}

func (h *ContactHandler) GetContacts(c *gin.Context) {
	c.JSON(http.StatusOK, h.contactUsecase.GetContacts())
}
