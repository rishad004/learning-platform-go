package pkg

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID uint
	Email  string
	jwt.StandardClaims
}

func JwtCreate(UserID uint, Email string) (string, error) {
	claims := Claims{
		UserID: UserID,
		Email:  Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your_secret_key"))

	return tokenString, err
}
