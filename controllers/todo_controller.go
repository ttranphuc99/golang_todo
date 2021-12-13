package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todoapi/config"
	"todoapi/database"
	"todoapi/dtos"
	"todoapi/models/constants"
	"todoapi/repository"
	"todoapi/services"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	GetAllTodo(c *gin.Context)
	InsertTodo(c *gin.Context)
	UpdateTodo(c *gin.Context)
	GetTodoByID(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type TodoControllerStruct struct {
	service services.TodoService
	config  config.Config
}

func NewTodoController(config config.Config) *TodoControllerStruct {
	return &TodoControllerStruct{
		config: config,
	}
}

// init
func (controller *TodoControllerStruct) init() error {
	dbHandler := database.NewDatabaseStruct(controller.config)
	repository, error := repository.NewTodoRepository(dbHandler, controller.config)

	if error != nil {
		log.Println(error)
		return error
	}

	controller.service = services.NewTodoServiceStruct(repository, controller.config)
	return nil
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
					ErrorMessage: controller.config.InvalidTodoStatusArgument,
				},
			)
			return
		}

		status = parsedStatus
	}

	currentUserId := c.GetString(controller.config.TokenCurrentUserId)
	currentUserRole := int(c.GetFloat64(controller.config.TokenCurrentUserRole))

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

	todoDTO.OwnerId = c.GetString(controller.config.TokenCurrentUserId)

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

	newTodo.OwnerId = c.GetString(controller.config.TokenCurrentUserId)

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

	owner := c.GetString(controller.config.TokenCurrentUserId)
	currentUserRole := int(c.GetFloat64(controller.config.TokenCurrentUserRole))

	var todo dtos.TodoDTO
	if currentUserRole == constants.RoleAdmin {
		todo, err = controller.service.GetTodoByID(id)
	} else {
		todo, err = controller.service.GetTodoByIDAndOwner(id, owner)
	}

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

// delete todo
func (controller *TodoControllerStruct) DeleteTodo(c *gin.Context) {
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

	owner := c.GetString(controller.config.TokenCurrentUserId)
	err = controller.service.DeleteTodo(id, owner)

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

	handleSuccess(c, dtos.MessageDTO{Message: "Action success."})
}
