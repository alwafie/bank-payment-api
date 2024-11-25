package utils

import (
	"belajar/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secretkey123")

type CustomClaims struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(customer *models.Customer) (string, error) {
	claims := CustomClaims{
		customer.ID,
		customer.Name,
		customer.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretKey)

	return ss, err
}

func ValidateToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok || !token.Valid {
		return nil, err
	}

	if IsTokenBlacklisted(tokenString) {
		return nil, err
	}
	return claims, nil
}
