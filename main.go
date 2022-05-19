package main

import (
	"log"
	"user_auth/controller"
	"user_auth/db"
	"user_auth/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.NewDB()
	e := controller.NewController(db)

	server := gin.Default()
	server.POST("/signup", e.Signup)
	server.POST("/login", e.Login)
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	server.Use(middleware.AuthorizeJWT())
	server.GET("/get", e.EmployeeList)

	log.Fatal(server.Run(":8080"))
}
