package pkg

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	AdminID uint
	Email   string
	jwt.StandardClaims
}

func JwtCreate(AdminID uint, Email string) (string, error) {
	claims := Claims{
		AdminID: AdminID,
		Email:   Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your_secret_key"))

	return tokenString, err
}
