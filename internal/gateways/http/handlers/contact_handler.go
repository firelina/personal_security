package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_security/internal/domain"
	"personal_security/internal/gateways/http/models"
	"personal_security/internal/usecase"
	"strconv"
)

type ContactHandler struct {
	contactUsecase *usecase.ContactService
}

func NewContactHandler(u *usecase.ContactService) *ContactHandler {
	return &ContactHandler{u}
}

func (h *ContactHandler) CreateEvent(c *gin.Context) {
	var newContact models.CreateContactRequest
	if err := c.ShouldBindJSON(&newContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	contactID, err := h.contactUsecase.CreateContact(ctx, &domain.Contact{
		UserID: newContact.UserID,
		Email:  newContact.Email,
		Phone:  newContact.Phone,
		Name:   newContact.Name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusCreated, contactID)
}

func (h *ContactHandler) GetContacts(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	contacts, err := h.contactUsecase.GetContacts(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, models.NewContactResponse(contacts))
}
