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
	var (
		userData         viewmodels.CreateUserViewmodel
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

	returnedUser, err := ctrl.service.CreateUserService(&userEntity)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	err = structMapper.Decode(returnedUser, &userResponseData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// userResponseData.Name = returnedUser.Name
	// userResponseData.Email = returnedUser.Email
	// userResponseData.CreatedAt = returnedUser.CreatedAt
	// userResponseData.UpdatedAt = returnedUser.UpdatedAt

	c.JSON(http.StatusOK, userResponseData)
}
