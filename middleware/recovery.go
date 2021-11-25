package middleware

import (
	"log"
	"net/http"
	"runtime"
	"runtime/debug"
	"todoapi/dtos"

	"github.com/gin-gonic/gin"
)

func Recover() gin.RecoveryFunc {
	return func(c *gin.Context, recovered interface{}) {
		msg := ""

		switch e := recovered.(type) {
		case string:
			msg = "recovered (string) panic:" + e
		case runtime.Error:
			msg = "recovered (runtime.Error) panic:" + e.Error()
		case error:
			msg = "recovered (error) panic:" + e.Error()
		default:
			msg = "internal error"
		}
		log.Println(msg)
		log.Println(string(debug.Stack()))
		c.JSON(http.StatusInternalServerError, dtos.BadRequestResponse{ErrorCode: "", ErrorMessage: msg})
	}
}
