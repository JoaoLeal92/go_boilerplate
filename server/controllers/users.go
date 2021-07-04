package controllers

import (
	"net/http"

	"github.com/JoaoLeal92/go_boilerplate/domain/contract"
	"github.com/JoaoLeal92/go_boilerplate/infra/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type createUserData struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// UsersController user controller
type UsersController struct {
	service contract.UserService
	cfg     *config.Config
}

// NewUsersController instantiates a new controller
func NewUsersController(s contract.UserService, cfg *config.Config) *UsersController {
	return &UsersController{
		service: s,
		cfg:     cfg,
	}
}

// CreateUserController creates new user
func (ctrl *UsersController) CreateUserController(c *gin.Context) {
	// Validate input
	var userData createUserData
	if err := c.ShouldBindBodyWith(&userData, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	returnedUser, err := ctrl.service.CreateUserService(userData.Name, userData.Email, userData.Password)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, returnedUser)
}
