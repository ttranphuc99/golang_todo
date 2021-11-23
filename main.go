package main

import (
	"log"
	"todoapi/controllers"
	"todoapi/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	todoController, accountController := initApp()

	todoRoutes := router.Group("/todo")
	{
		todoRoutes.GET("/", middleware.CheckToken(), todoController.GetAllTodo)
		todoRoutes.POST("/", middleware.CheckToken(), todoController.InsertTodo)
		todoRoutes.PUT("/", todoController.UpdateTodo)
		todoRoutes.GET("/:id", todoController.GetTodoByID)
	}

	router.POST("/login", accountController.Login)

	router.Run("localhost:9000")
}

func initApp() (todoController controllers.TodoController, accountController controllers.AccountController) {
	todoController = &controllers.TodoControllerStruct{}
	accountController = &controllers.AccountControllerStruct{}

	error := todoController.Init()

	if error != nil {
		log.Panicln(error)
		panic(error)
	}

	error = accountController.Init()

	if error != nil {
		log.Panicln(error)
		panic(error)
	}

	return todoController, accountController
}
