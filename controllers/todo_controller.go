package controllers

import (
	"log"
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
	Init() error
}

type TodoControllerStruct struct {
	service services.TodoService
}

// init
func (controller *TodoControllerStruct) Init() error {
	controller.service = &services.TodoServiceStruct{}
	return controller.service.Init()
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
		result, err := controller.service.GetAllTodo(status)

		if err != nil {
			handleBadRequest(
				c,
				dtos.BadRequestResponse{
					ErrorMessage: err.Error(),
				},
			)
			return
		}
		handleSuccess(c, result)
		return
	}

	// get all to do
	result, err := controller.service.GetAllTodo(constants.TodoStatusAll)

	if err != nil {
		log.Panicln(err)
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}
	handleSuccess(c, result)
}

// InsertTodo
func (controller *TodoControllerStruct) InsertTodo(c *gin.Context) {
	newTodo := &models.Todo{}

	if err := c.BindJSON(newTodo); err != nil {
		log.Panicln(err)
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	newTodo.OwnerId = c.GetString(config.TOKEN_CURRENT_USER_ID)

	resultTodo, err := controller.service.InsertTodo(newTodo)

	if err != nil {
		log.Panicln(err)
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	c.IndentedJSON(http.StatusCreated, resultTodo)
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
