package controllers

import (
	"log"
	"todoapi/dtos"
	"todoapi/models"
	"todoapi/services"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	Login(c *gin.Context)
}

type AccountControllerStruct struct {
	service services.UserAccountService
}

func (controller *AccountControllerStruct) init() error {
	tempService := &services.UserAccountServiceStruct{}
	controller.service = tempService
	return controller.service.Init()
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
