package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_security/internal/domain"
	"personal_security/internal/gateways/http/models"
	"personal_security/internal/usecase"
)

type UserHandler struct {
	userUsecase *usecase.UserService
}

func NewUserHandler(u *usecase.UserService) *UserHandler {
	return &UserHandler{u}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var newUser *models.CreateUserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	userID, err := h.userUsecase.RegisterUser(ctx, &domain.User{
		Email:    newUser.Email,
		Name:     newUser.Name,
		Password: newUser.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	c.JSON(http.StatusCreated, userID)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var loginUser models.LoginUserRequest
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	user, err := h.userUsecase.LoginUser(ctx, &domain.User{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user.Name})

}
