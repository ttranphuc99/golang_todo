package services

import (
	"net/http"
	"strconv"
	"todoapi/models"
	"todoapi/models/constants"

	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{
	{ID: 1, Content: "11111", Status: constants.StatusActive},
	{ID: 2, Content: "2222", Status: constants.StatusCompleted},
	{ID: 3, Content: "3333", Status: constants.StatusActive},
}

func GetAllTodo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func InsertTodo(c *gin.Context) {
	var newTodo models.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodoByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found todo with id = " + c.Param("id")})
		return
	}

	for _, todo := range todos {
		if id == todo.ID {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found todo with id = " + c.Param("id")})
}
