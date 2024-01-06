package service

import (
	"fmt"
	"os"

	"userauthservice/model"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct{}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (s *JwtService) GenerateJwt(user *model.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})
	secretKey := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error signing the JWT")
	}
	return tokenStr
}

func (s *JwtService) ValidateJwt(tokenStr string) (bool, string) {
	type CustomClaims struct {
		username string
		jwt.RegisteredClaims
	}
	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("Token parsing error:", err)
		return false, ""
	}
	if !token.Valid {
		fmt.Println("Invalid token")
		return false, ""
	}
	if claims, ok := token.Claims.(*CustomClaims); ok {
		return true, claims.username
	}
	return false, ""
}
