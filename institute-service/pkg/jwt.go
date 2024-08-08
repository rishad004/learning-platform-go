package pkg

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func Middleware(c *gin.Context) {
	tokenString, err := c.Cookie("Jwt-Admin")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
	claims := &Claims{}
	token, er := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if er != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	c.Set("id", claims.AdminID)
	c.Next()
}
