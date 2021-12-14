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

	todoRoutes := router.Group("/todo").Use(middleware.CheckToken(config))
	{
		todoRoutes.GET("/", todoController.GetAllTodo)
		todoRoutes.POST("/", todoController.InsertTodo)
		todoRoutes.PUT("/", todoController.UpdateTodo)
		todoRoutes.GET("/:id", todoController.GetTodoByID)
		todoRoutes.DELETE("/:id", todoController.DeleteTodo)
	}

	router.POST("/login", accountController.Login)

	router.Run("localhost:9000")
}
