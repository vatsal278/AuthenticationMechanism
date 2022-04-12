package

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"net/http"
	"user_auth/service"
)
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.JWTAuthService().ValidateToken(tokenString)

		if token.Valid {
			claims:= token.Claims.(jwt.MapClaims)
			fmt.Println((claims))
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}