package main

import (
	"log"
	"user_auth/controller"
	"user_auth/db"
	"user_auth/middleware"
	"user_auth/service"

	"github.com/gin-gonic/gin"
)

func main() {
	var loginService = service.StaticLoginService()
	var jwtService = service.JWTAuthService()
	db := db.NewDB()
	e := controller.NewController(db, loginService, jwtService)

	server := gin.Default()
	server.POST("/signup", e.Signup)
	server.POST("/login", e.Login)
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	server.Use(middleware.AuthorizeJWT())
	server.GET("/get", e.EmployeeList)

	log.Fatal(server.Run(":9090"))
}
