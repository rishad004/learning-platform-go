package pkg

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID uint
	Email  string
	jwt.StandardClaims
}

func Middleware(c *gin.Context) {
	tokenString, err := c.Cookie("Jwt-User")

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

	c.Set("id", claims.UserID)
	c.Next()
}
