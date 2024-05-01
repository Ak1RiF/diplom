package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/todoApp/pkg/models"
)

var (
	SigningKey = []byte("secret generating key")
)

// Custom claims
type CustomClaims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(user models.User) (string, error) {
	claims := CustomClaims{
		user.Id,
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
