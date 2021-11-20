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

type TodoController interface {
	GetAllTodo(c *gin.Context)
	InsertTodo(c *gin.Context)
	UpdateTodo(c *gin.Context)
	GetTodoByID(c *gin.Context)
	Init()
}

type TodoControllerStruct struct {
	service services.TodoService
}

// init
func (controller *TodoControllerStruct) Init() {
	controller.service = &services.TodoServiceStruct{}
	controller.service.Init()
}

// GetAllTodo
func (controller *TodoControllerStruct) GetAllTodo(c *gin.Context) {
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
		result := controller.service.GetAllTodo(status)
		handleSuccess(c, result)
		return
	}

	// get all to do
	result := controller.service.GetAllTodo(constants.TodoStatusAll)
	handleSuccess(c, result)
}

// InsertTodo
func (controller *TodoControllerStruct) InsertTodo(c *gin.Context) {
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

	newTodo, err := controller.service.InsertTodo(newTodo)

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
func (controller *TodoControllerStruct) UpdateTodo(c *gin.Context) {
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

	newTodo, err := controller.service.UpdateTodo(newTodo)

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
func (controller *TodoControllerStruct) GetTodoByID(c *gin.Context) {
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

	todo, err := controller.service.GetTodoByID(id)

	if err != nil {
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

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
