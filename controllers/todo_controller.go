package controllers

import (
	"net/http"
	"strconv"
	"todoapi/config"
	"todoapi/dtos"
	"todoapi/models"
	"todoapi/models/constants"
	"todoapi/services"

	"github.com/gin-gonic/gin"
)

// GetAllTodo
func GetAllTodo(c *gin.Context) {
	// get status filter
	statusFilter := c.Query("status")

	// status filter has value
	if statusFilter != "" {
		// parse value to int
		status, err := strconv.ParseInt(statusFilter, 10, 64)

		// parse value failed
		if err != nil {
			handleBadRequest(
				c,
				dtos.BadRequestResponse{
					ErrorMessage: config.InvalidTodoStatusArgument,
				},
			)
			return
		}

		// parse value success
		result := services.GetAllTodo(status)
		handleSuccess(c, result)
		return
	}

	// get all to do
	result := services.GetAllTodo(constants.TodoStatusAll)
	handleSuccess(c, result)
}

// InsertTodo
func InsertTodo(c *gin.Context) {
	var newTodo models.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	newTodo, err := services.InsertTodo(newTodo)

	if err != nil {
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	c.IndentedJSON(http.StatusCreated, newTodo)
}

// UpdateTodo
func UpdateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	newTodo, err := services.UpdateTodo(newTodo)

	if err != nil {
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	handleSuccess(c, newTodo)
}

// GetTodoByID
func GetTodoByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	todo, err := services.GetTodoByID(id)

	handleSuccess(c, todo)
}

// handle bad request
func handleBadRequest(c *gin.Context, errorResponse dtos.BadRequestResponse) {
	c.IndentedJSON(http.StatusBadRequest, errorResponse)
}

// handle success
func handleSuccess(c *gin.Context, data interface{}) {
	c.IndentedJSON(http.StatusOK, data)
}
