package controllers

import (
	"net/http"
	"todoapi/dtos"

	"github.com/gin-gonic/gin"
)

// handle bad request
func handleBadRequest(c *gin.Context, errorResponse dtos.BadRequestResponse) {
	c.IndentedJSON(http.StatusBadRequest, errorResponse)
}

// handle success
func handleSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func handlerError(c *gin.Context, errorResponse dtos.BadRequestResponse) {
	c.IndentedJSON(http.StatusInternalServerError, errorResponse)
}
