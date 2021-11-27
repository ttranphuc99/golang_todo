package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todoapi/config"
	"todoapi/dtos"
	"todoapi/models/constants"
	"todoapi/services"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	GetAllTodo(c *gin.Context)
	InsertTodo(c *gin.Context)
	UpdateTodo(c *gin.Context)
	GetTodoByID(c *gin.Context)
}

type TodoControllerStruct struct {
	service services.TodoService
}

// init
func (controller *TodoControllerStruct) init() error {
	controller.service = &services.TodoServiceStruct{}
	return controller.service.Init()
}

// GetAllTodo
func (controller *TodoControllerStruct) GetAllTodo(c *gin.Context) {
	if error := controller.init(); error != nil {
		log.Println(error)
		return
	}

	// get status filter
	statusFilter := c.Query("status")
	status := constants.TodoStatusAll

	// status filter has value
	if statusFilter != "" {
		// parse value to int
		parsedStatus, err := strconv.Atoi(statusFilter)

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

		status = parsedStatus
	}

	currentUserId := c.GetString(config.TOKEN_CURRENT_USER_ID)
	currentUserRole := int(c.GetFloat64(config.TOKEN_CURRENT_USER_ROLE))

	// get all to do
	result, err := controller.service.GetAllTodo(status, currentUserId, currentUserRole)

	if err != nil {
		log.Println(err)
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
	if error := controller.init(); error != nil {
		log.Println(error)
		return
	}

	todoDTO := &dtos.TodoDTO{}

	if err := c.BindJSON(todoDTO); err != nil {
		log.Println(err)
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	todoDTO.OwnerId = c.GetString(config.TOKEN_CURRENT_USER_ID)

	resultTodo, err := controller.service.InsertTodo(todoDTO)

	if err != nil {
		log.Println(err)
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusCreated, resultTodo)
}

// UpdateTodo
func (controller *TodoControllerStruct) UpdateTodo(c *gin.Context) {
	if error := controller.init(); error != nil {
		log.Println(error)
		return
	}

	var newTodo dtos.TodoDTO
	if err := c.BindJSON(&newTodo); err != nil {
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	newTodo.OwnerId = c.GetString(config.TOKEN_CURRENT_USER_ID)

	updatedTodo, err := controller.service.UpdateTodo(newTodo)

	if err != nil {
		if err == sql.ErrNoRows {
			handleNotFound(
				c,
				dtos.BadRequestResponse{
					ErrorMessage: "Not found todo with ID " + fmt.Sprintf("%d", newTodo.ID) + " of " + newTodo.OwnerId,
				})
			return
		}
		handleBadRequest(
			c,
			dtos.BadRequestResponse{
				ErrorMessage: err.Error(),
			},
		)
		return
	}

	handleSuccess(c, updatedTodo)
}

// GetTodoByID
func (controller *TodoControllerStruct) GetTodoByID(c *gin.Context) {
	if error := controller.init(); error != nil {
		log.Println(error)
		return
	}

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

	owner := c.GetString(config.TOKEN_CURRENT_USER_ID)
	todo, err := controller.service.GetTodoByID(id, owner)

	if err != nil {
		if err == sql.ErrNoRows {
			handleNotFound(
				c,
				dtos.BadRequestResponse{
					ErrorMessage: "Not found todo with ID " + fmt.Sprintf("%d", id) + " of " + owner,
				})
			return
		}
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
