package services

import (
	"errors"
	"net/http"
	"strconv"
	"todoapi/models"
	"todoapi/models/constants"

	"github.com/gin-gonic/gin"
)

var todos = []*models.Todo{
	{ID: 1, Content: "11111", Status: constants.StatusActive},
	{ID: 2, Content: "2222", Status: constants.StatusCompleted},
	{ID: 3, Content: "3333", Status: constants.StatusActive},
}

func GetAllTodo(c *gin.Context) {
	statusFilter := c.Query("status")

	if statusFilter != "" {
		status, err := strconv.ParseInt(statusFilter, 10, 64)

		if err == nil {
			todosRes := []models.Todo{}

			for _, todo := range todos {
				if status == int64(todo.Status) {
					todosRes = append(todosRes, *todo)
				}
			}

			c.IndentedJSON(http.StatusOK, todosRes)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func InsertTodo(c *gin.Context) {
	var newTodo models.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	_, err := findById(newTodo.ID)

	if err == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Duplicate ID"})
		return
	}

	todos = append(todos, &newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodoByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found todo with id = " + c.Param("id")})
		return
	}

	todo, err := findById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found todo with id = " + c.Param("id")})
	} else {
		c.IndentedJSON(http.StatusOK, *todo)
	}
}

func UpdateTodo(c *gin.Context) {
	var newTodo models.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	oldTodo, err := findById(newTodo.ID)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found todo with id = " + c.Param("id")})
	} else {
		oldTodo.Content = newTodo.Content
		oldTodo.Status = newTodo.Status

		c.IndentedJSON(http.StatusOK, *oldTodo)
	}
}

func findById(id int64) (todo *models.Todo, e error) {
	for idx, todo := range todos {
		if id == todo.ID {
			return todos[idx], nil
		}
	}
	return &models.Todo{}, errors.New("not found")
}
