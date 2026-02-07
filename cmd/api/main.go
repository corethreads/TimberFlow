package main

import (
	"server/config"
	"server/internal/auth/controller"
	"server/internal/auth/repository"
	"server/internal/auth/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadDB()
	config.RunMigrations()

	//setup Layers
	//repository
	userRepo := repository.NewUserRepository(config.Database)

	//services
	userService := service.NewauthService(userRepo)

	//controllers
	userController := controller.NewAuthController(userService)

	router := gin.Default()

	//Cors middleware setup
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/register", userController.Register)
	router.Run("0.0.0.0:8080")

}
