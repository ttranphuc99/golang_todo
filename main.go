package main

import (
	"todoapi/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todo", services.GetAllTodo)
	router.POST("/todo", services.InsertTodo)
	router.PUT("/todo", services.UpdateTodo)
	router.GET("/todo/:id", services.GetTodoByID)

	router.Run("localhost:9000")
}
