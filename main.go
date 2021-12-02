package main

import (
	"todoapi/controllers"
	"todoapi/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.CustomRecovery(middleware.Recover()))
	todoController := controllers.TodoControllerStruct{}
	accountController := controllers.AccountControllerStruct{}

	todoRoutes := router.Group("/todo")
	{
		todoRoutes.GET("/", middleware.CheckToken(), todoController.GetAllTodo)
		todoRoutes.POST("/", middleware.CheckToken(), todoController.InsertTodo)
		todoRoutes.PUT("/", middleware.CheckToken(), todoController.UpdateTodo)
		todoRoutes.GET("/:id", middleware.CheckToken(), todoController.GetTodoByID)
		todoRoutes.DELETE("/:id", middleware.CheckToken(), todoController.DeleteTodo)
	}

	router.POST("/login", accountController.Login)

	router.Run("localhost:9000")
}
