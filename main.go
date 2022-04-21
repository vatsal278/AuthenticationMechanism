package main

import (
	"fmt"
	"log"
	"net/http"
	"user_auth/controller"
	"user_auth/middleware"
	"user_auth/service"

	"github.com/gin-gonic/gin"
)

func main() {
	var loginService = service.StaticLoginService()
	var jwtService = service.JWTAuthService()
	var loginController = controller.LoginHandler(loginService, jwtService)

	server := gin.Default()
	server.POST("/signup", func(ctx *gin.Context) {
		loginController.Signup(ctx)
	})
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		fmt.Print(token)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	server.Use(middleware.AuthorizeJWT())
	server.GET("/get", func(ctx *gin.Context) {

		_ = loginController.EmployeeList(ctx)
	})

	log.Fatal(server.Run(":8080"))
}
