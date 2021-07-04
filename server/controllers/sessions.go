package controllers

import (
	"net/http"

	"github.com/JoaoLeal92/go_boilerplate/domain/contract"
	"github.com/JoaoLeal92/go_boilerplate/domain/entities"
	"github.com/JoaoLeal92/go_boilerplate/infra/config"
	"github.com/JoaoLeal92/go_boilerplate/infra/mapper"
	"github.com/JoaoLeal92/go_boilerplate/server/viewmodels"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// SessionsController session controller struct
type SessionsController struct {
	service contract.SessionService
	cfg     *config.Config
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

	var (
		userData         viewmodels.CreateSessionViewModel
		userResponseData viewmodels.UserResponseViewmodel
		userEntity       entities.User
	)

	structMapper := mapper.NewMapper(true)

	if err := c.ShouldBindBodyWith(&userData, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := structMapper.Decode(&userData, &userEntity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, user, err := ctrl.service.AuthenticateUserService(&userEntity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = structMapper.Decode(user, &userResponseData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "user": userResponseData})
	return
}
