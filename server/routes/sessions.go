package routes

import (
	"github.com/JoaoLeal92/goals_backend/server/controllers"
	"github.com/gin-gonic/gin"
)

// SessionRoutes users route group
func SessionRoutes(router *gin.Engine, controller *controllers.SessionsController) {
	sessionsRoutes := router.Group("/sessions")
	{
		sessionsRoutes.POST("/", controller.CreateSession)
	}
}
