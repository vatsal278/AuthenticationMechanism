package service

import {
	"https://github.com/dgrijalva/jwt-go"
}
type JWTService interface {
	GenerateToken(emai string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Name string 'jsonname'
	User bool  'json' : "user"
	jwt.StandardClaims
}

type jwtService struct {
	secretkey string
	issure string
}

func JWTAuthService() JWTService{
	return &jwtService{
		secretkey: getSecretKey(),
		issure: "vatsal"
	}
}
