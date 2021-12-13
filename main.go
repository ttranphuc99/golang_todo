package main

import (
	"log"
	"todoapi/config"
	"todoapi/controllers"
	"todoapi/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Error calling config ", err)
	}

	router := gin.Default()
	router.Use(gin.CustomRecovery(middleware.Recover()))
	todoController := controllers.NewTodoController(config)
	accountController := controllers.NewAccountController(config)

	todoRoutes := router.Group("/todo")
	{
		todoRoutes.GET("/", middleware.CheckToken(config), todoController.GetAllTodo)
		todoRoutes.POST("/", middleware.CheckToken(config), todoController.InsertTodo)
		todoRoutes.PUT("/", middleware.CheckToken(config), todoController.UpdateTodo)
		todoRoutes.GET("/:id", middleware.CheckToken(config), todoController.GetTodoByID)
		todoRoutes.DELETE("/:id", middleware.CheckToken(config), todoController.DeleteTodo)
	}

	router.POST("/login", accountController.Login)

	router.Run("localhost:9000")
}
