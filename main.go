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
		todoRoutes.GET("/", middleware.CheckToken(), todoController.GetAllTodo, gin.CustomRecovery(middleware.Recover()))
		todoRoutes.POST("/", middleware.CheckToken(), todoController.InsertTodo, gin.CustomRecovery(middleware.Recover()))
		todoRoutes.PUT("/", middleware.CheckToken(), todoController.UpdateTodo, gin.CustomRecovery(middleware.Recover()))
		todoRoutes.GET("/:id", middleware.CheckToken(), todoController.GetTodoByID, gin.CustomRecovery(middleware.Recover()))
	}

	router.POST("/login", accountController.Login)

	router.Run("localhost:9000")
}
