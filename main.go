package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/JoaoLeal92/go_boilerplate/domain/services"
	"github.com/JoaoLeal92/go_boilerplate/domain/services/sessions"
	"github.com/JoaoLeal92/go_boilerplate/domain/services/user"
	"github.com/JoaoLeal92/go_boilerplate/infra/config"
	"github.com/JoaoLeal92/go_boilerplate/infra/hash"
	"github.com/JoaoLeal92/go_boilerplate/repositories"
	"github.com/JoaoLeal92/go_boilerplate/server/controllers"
	"github.com/JoaoLeal92/go_boilerplate/server/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	conn, err := repositories.ConnectDataBase(cfg.Db)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	hashProvider := hash.NewProvider()
	svc := services.NewService(conn, hashProvider, &cfg)

	var (
		userService    = user.NewUserService(svc)
		sessionService = sessions.NewSessionService(svc)

		usersConroller     = controllers.NewUsersController(userService, &cfg)
		sessionsController = controllers.NewSessionsController(sessionService, &cfg)
	)

	r := gin.Default()

	routes.UsersRoutes(r, usersConroller)
	routes.SessionRoutes(r, sessionsController)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", hello)

	r.Run(":3000")
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "API working"})
}
