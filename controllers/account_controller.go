package controllers

import (
	"log"
	"todoapi/config"
	"todoapi/database"
	"todoapi/dtos"
	"todoapi/models"
	"todoapi/repository"
	"todoapi/services"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	Login(c *gin.Context)
}

type AccountControllerStruct struct {
	service services.UserAccountService
	config  config.Config
}

func NewAccountController(config config.Config) *AccountControllerStruct {
	controller := &AccountControllerStruct{
		config: config,
	}
	return controller
}

func (controller *AccountControllerStruct) init() error {
	dbHandler := database.NewDatabaseStruct(controller.config)
	repository, error := repository.NewUserAccountRepository(dbHandler, controller.config)

	if error != nil {
		log.Println("Error " + error.Error())
		return error
	}

	controller.service = services.NewUserAccountService(repository, controller.config)

	return nil
}

func (controller *AccountControllerStruct) Login(c *gin.Context) {
	if error := controller.init(); error != nil {
		log.Println(error)
		return
	}
	var user models.UserAccount

	error := c.BindJSON(&user)

	if error != nil {
		log.Println(error)
		handleBadRequest(c, dtos.BadRequestResponse{
			ErrorMessage: error.Error(),
		})
	}

	resultUser, error := controller.service.Login(user)

	if error != nil {
		log.Println(error)
		handleError(c, dtos.BadRequestResponse{
			ErrorMessage: error.Error(),
		})
		return
	}

	handleSuccess(c, resultUser)
}
