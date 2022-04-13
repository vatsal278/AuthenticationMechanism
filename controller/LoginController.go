package controller

import (
	"user_auth/Credentials"
	"user_auth/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

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

	isUserAuthenticated := controller.loginService.LogInUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jwtService.GenerateToken(credential.Email, true)
	}
	return ""
}
