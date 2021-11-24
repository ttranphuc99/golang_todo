package main

import (
	"todoapi/controllers"
	"todoapi/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	todoController := controllers.TodoControllerStruct{}
	accountController := controllers.AccountControllerStruct{}

	todoRoutes := router.Group("/todo")
	{
		todoRoutes.GET("/", middleware.CheckToken(), todoController.GetAllTodo)
		todoRoutes.POST("/", middleware.CheckToken(), todoController.InsertTodo)
		todoRoutes.PUT("/", middleware.CheckToken(), todoController.UpdateTodo)
		todoRoutes.GET("/:id", middleware.CheckToken(), todoController.GetTodoByID)
	}

	router.POST("/login", accountController.Login)

	router.Run("localhost:9000")
}
