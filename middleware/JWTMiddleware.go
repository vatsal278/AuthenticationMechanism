package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"user_auth/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware is working")
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		newHeader := strings.Split(authHeader, " ")
		if len(newHeader) != 2 || newHeader[0] != BEARER_SCHEMA || len(strings.Trim(newHeader[1], " ")) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := service.JWTAuthService().ValidateToken(strings.Trim(newHeader[1], " "))
		if err != nil {
			log.Print(err, "compared literals are not same")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println((claims))
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
