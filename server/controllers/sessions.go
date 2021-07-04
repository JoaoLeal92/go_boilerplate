package controllers

import (
	"net/http"

	"github.com/JoaoLeal92/go_boilerplate/domain/contract"
	"github.com/JoaoLeal92/go_boilerplate/infra/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// SessionsController session controller struct
type SessionsController struct {
	service contract.SessionService
	cfg     *config.Config
}

type authenticateUserData struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// NewSessionsController creates new session controller
func NewSessionsController(s contract.SessionService, cfg *config.Config) *SessionsController {
	return &SessionsController{
		service: s,
		cfg:     cfg,
	}
}

// CreateSession creates session for registered user
func (ctrl *SessionsController) CreateSession(c *gin.Context) {
	// Validate body content
	var userData authenticateUserData
	if err := c.ShouldBindBodyWith(&userData, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Calls service for authentication
	tokenString, user, err := ctrl.service.AuthenticateUserService(userData.Email, userData.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "user": user})
	return
}
