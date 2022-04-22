package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(emai string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
	DecodeToken(token string) (jwt.MapClaims, error)
}

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issure    string
}

func JWTAuthService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issure:    "vatsal",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtService) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 100000).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			err := fmt.Errorf("Invalid token", token.Header["alg"])
			return nil, err
		}
		return []byte(service.secretKey), nil
	})

}

func (service *jwtService) DecodeToken(encodedToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return token, nil
	})
	claims := token.Claims.(jwt.MapClaims)
	return claims, err
}
