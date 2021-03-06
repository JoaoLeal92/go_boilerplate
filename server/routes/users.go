package routes

import (
	"github.com/JoaoLeal92/go_boilerplate/server/controllers"
	"github.com/gin-gonic/gin"
)

// UsersRoutes users route group
func UsersRoutes(router *gin.Engine, controller *controllers.UsersController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", controller.CreateUserController)
	}
}
