package middleware

import (
	"log"
	"net/http"
	"strings"
	"todoapi/config"
	"todoapi/dtos"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckToken(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		if header == "" {
			handleUnauthorized(c, "Invalid Token")
			return
		}

		splitToken := strings.Split(header, "Bearer ")

		if len(splitToken) != 2 {
			handleUnauthorized(c, "Invalid Token")
			return
		}

		reqToken := splitToken[1]

		if reqToken == "" {
			handleUnauthorized(c, "Invalid Token")
			return
		}

		token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.SecretKeyJwt), nil
		})

		if err != nil {
			handleUnauthorized(c, err.Error())
		} else if !token.Valid {
			log.Printf("Token is invalid %s\n", reqToken)
			handleUnauthorized(c, "Invalid Token")
		}

		claims, isOk := token.Claims.(jwt.MapClaims)

		if isOk {
			c.Set(config.TokenCurrentUserId, claims[config.TokenCurrentUserId])
			c.Set(config.TokenCurrentUserRole, claims[config.TokenCurrentUserRole])
		} else {
			log.Panicf("Cannot extract claims from token %s\n", reqToken)
			handleUnauthorized(c, "Invalid Token")
		}

	}
}

func handleUnauthorized(c *gin.Context, errorMsg string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, dtos.BadRequestResponse{ErrorMessage: errorMsg})
}

// func handleForbidden(c *gin.Context) {
// 	c.IndentedJSON(http.StatusForbidden, nil)
// }
