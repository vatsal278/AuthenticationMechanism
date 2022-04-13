package controller

import(
	"github.com/gin-gonic/gin"
	"user_auth/credentials"
	"user_auth/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService service.JWTService
}

func LoginHandler(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController {
		loginService : loginService
		jwtService : jwtService
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential credentials.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}

	isUserAuthenticated := controller.loginService.LogInUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jwtService.GenerateToken(credential.Email, true )
	}
	return ""
}