package controllers

import (
	"log"
	"todoapi/dtos"
	"todoapi/models"
	"todoapi/services"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	Init() error
	Login(c *gin.Context)
}

type AccountControllerStruct struct {
	service services.UserAccountService
}

func (controller *AccountControllerStruct) Init() error {
	tempService := &services.UserAccountServiceStruct{}
	controller.service = tempService
	return controller.service.Init()
}

func (controller *AccountControllerStruct) Login(c *gin.Context) {
	var user models.UserAccount

	error := c.BindJSON(&user)

	if error != nil {
		log.Panicln(error)
		handleBadRequest(c, dtos.BadRequestResponse{
			ErrorMessage: error.Error(),
		})
	}

	resultUser, error := controller.service.Login(user)

	if error != nil {
		log.Panicln(error)
		handlerError(c, dtos.BadRequestResponse{
			ErrorMessage: error.Error(),
		})
	}

	handleSuccess(c, resultUser)
}
