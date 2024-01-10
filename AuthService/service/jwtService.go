package service

import (
	"fmt"
	"os"

	"authservice/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtService struct{}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (s *JwtService) GenerateJwt(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	secretKey := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error signing the JWT")
	}
	return tokenStr, err
}

func (s *JwtService) ValidateJwt(tokenStr string) uuid.UUID {
	type CustomClaims struct {
		ID uuid.UUID
		jwt.RegisteredClaims
	}
	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("Token parsing error:", err)
		return uuid.Nil
	}
	if !token.Valid {
		fmt.Println("Invalid token")
		return uuid.Nil
	}
	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims.ID
	}
	return uuid.Nil
}
