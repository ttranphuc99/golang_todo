package main

import (
	"todoapi/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todo", controllers.GetAllTodo)
	router.POST("/todo", controllers.InsertTodo)
	router.PUT("/todo", controllers.UpdateTodo)
	router.GET("/todo/:id", controllers.GetTodoByID)

	router.Run("localhost:9000")
}
