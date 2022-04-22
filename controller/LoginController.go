package controller

import (
	"fmt"
	"log"
	"strings"
	"user_auth/Credentials"
	"user_auth/helpers"
	"user_auth/models"
	"user_auth/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
	EmployeeList(ctx *gin.Context) string
	Signup(ctx *gin.Context)
}

type Employee struct {
	Id   string `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	City string `form:"city" json:"city"`
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

	if ctx.BindJSON(&credential) != nil {
		ctx.JSON(406, gin.H{"message": "Provide required details"})
		ctx.Abort()
		return ""
	}

	result, err := userModel.GetUserByEmail(credential.Email)

	if result.Email == "" {
		ctx.JSON(404, gin.H{"message": "User account was not found"})
		ctx.Abort()
		return ""
	}

	fmt.Println(result)

	if err != nil {
		ctx.JSON(400, gin.H{"message": "Problem logging into your account"})
		ctx.Abort()
		return ""
	}
	hashedPassword := []byte(result.Password)
	// Get the password provided in the request.body
	password := []byte(credential.Password)

	err = helpers.PasswordCompare(password, hashedPassword)

	if err != nil {
		ctx.JSON(403, gin.H{"message": "Invalid user credentials"})
		ctx.Abort()
		return ""
	}

	return controller.jwtService.GenerateToken(credential.Email, true)

}

func (controller *loginController) EmployeeList(ctx *gin.Context) string {
	// todo only loggedin user should get their name and email in the response

	authHeader := ctx.GetHeader("Authorization")
	newHeader := strings.Split(authHeader, " ")
	newHeader[0] = strings.Trim(newHeader[1], " ")
	claims, err := controller.jwtService.DecodeToken(newHeader[0])
	if err != nil {
		log.Print("Decode Unsuccessfull")
	}

	log.Printf("You are logged in as %s, Hope API is working Fine", claims["name"])
	return ""
}

var userModel = new(models.UserModel)

type UserController struct{}

func (controller *loginController) Signup(c *gin.Context) {
	var usercredential Credentials.SignUpCredentials

	if c.BindJSON(&usercredential) != nil {
		// specified response
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		// abort the request
		c.Abort()
		// return nothing
		return
	}

	result, _ := userModel.GetUserByEmail(usercredential.Email)

	if result.Email != "" {
		c.JSON(403, gin.H{"message": "Email is already in use"})
		c.Abort()
		return
	}

	err := userModel.Signup(usercredential)

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an account"})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"message": "New user account registered"})
}
