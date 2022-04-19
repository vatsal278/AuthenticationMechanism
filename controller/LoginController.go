package controller

import (
	"fmt"
	"user_auth/Credentials"
	"user_auth/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
	EmployeeList(ctx *gin.Context) string
}

type Employee struct {
	Id   string `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	City string `form:"city" json:"city"`
}

var employee Employee = Employee{"100", "vatsal", "jaipur"}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func LoginHandler(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential Credentials.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	fmt.Print(credential)
	isUserAuthenticated := controller.loginService.LogInUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jwtService.GenerateToken(credential.Email, true)
	}
	return ""
}

func (controller *loginController) EmployeeList(ctx *gin.Context) string {
	fmt.Printf("%s", employee)
	return ""
}
