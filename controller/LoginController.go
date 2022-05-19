package controller

import (
	"fmt"
	"log"
	"user_auth/Credentials"
	"user_auth/helpers"
	"user_auth/models"
	"user_auth/service"

	"github.com/gin-gonic/gin"
)

type ILoginController interface {
	Login(ctx *gin.Context)
	EmployeeList(ctx *gin.Context)
	Signup(ctx *gin.Context)
}

type Employee struct {
	Id   string `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	City string `form:"city" json:"city"`
}

type LoginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
	Db           models.IUserModel
}

func NewController(dbi models.IUserModel) ILoginController {
	return &LoginController{
		Db: dbi,
	}
}

func LoginHandler(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return LoginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (c *LoginController) Login(ctx *gin.Context) {
	var credential Credentials.LoginCredentials

	if ctx.BindJSON(&credential) != nil {
		ctx.JSON(406, gin.H{"message": "Provide required details"})
		ctx.Abort()
		return
	}

	result, err := c.Db.GetUserByEmail(credential.Email)

	if result.Email == "" {
		ctx.JSON(404, gin.H{"message": "User account was not found"})
		ctx.Abort()
		return
	}

	fmt.Println(result)

	if err != nil {
		ctx.JSON(400, gin.H{"message": "Problem logging into your account"})
		ctx.Abort()
		return
	}
	hashedPassword := []byte(result.Password)
	// Get the password provided in the request.body
	password := []byte(credential.Password)

	err = helpers.PasswordCompare(password, hashedPassword)

	if err != nil {
		ctx.JSON(403, gin.H{"message": "Invalid user credentials"})
		ctx.Abort()
		return
	}
	ctx.JSON(200, c.jwtService.GenerateToken(credential.Email, true))

}

func (c *LoginController) EmployeeList(ctx *gin.Context) {

	loggeduser, exist := ctx.Get("email")
	user, err := c.Db.GetUserByEmail(fmt.Sprint(loggeduser))
	if err != nil {
		log.Print("cant fetch user from db")
	}

	if !exist {
		log.Print("cannot pass variable accross middleware")
	}
	var usercredential struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	usercredential.Name = user.Name
	usercredential.Email = user.Email
	ctx.JSON(200, usercredential)

	log.Printf("You are logged in as %s, Hope API is working Fine", user.Name)
}

type UserController struct{}

func (c *LoginController) Signup(ctx *gin.Context) {
	var usercredential Credentials.SignUpCredentials

	if ctx.BindJSON(&usercredential) != nil {
		// specified response
		ctx.JSON(406, gin.H{"message": "Provide relevant fields"})
		// abort the request
		ctx.Abort()
		// return nothing
		return
	}

	result, _ := c.Db.GetUserByEmail(usercredential.Email)

	if result.Email != "" {
		ctx.JSON(403, gin.H{"message": "Email is already in use"})
		ctx.Abort()
		return
	}

	err := c.Db.Signup(usercredential)

	if err != nil {
		ctx.JSON(400, gin.H{"message": "Problem creating an account"})
		log.Println("Problem creating an account", err.Error())
		ctx.Abort()
		return
	}

	ctx.JSON(201, gin.H{"message": "New user account registered"})
}
