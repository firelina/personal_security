package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_security/internal/domain"
	"personal_security/internal/usecase"
)

type UserHandler struct {
	userUsecase *usecase.User
}

func NewUserHandler(u *usecase.User) *UserHandler {
	return &UserHandler{u}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var newUser *domain.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := h.userUsecase.RegisterUser(newUser)
	c.JSON(http.StatusCreated, userID)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var loginUser domain.User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUsecase.LoginUser(&loginUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})

}
