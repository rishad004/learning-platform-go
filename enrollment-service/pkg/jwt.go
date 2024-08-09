package pkg

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID uint
	Email  string
	jwt.StandardClaims
}

func Middleware(c *gin.Context) {
	tokenString, _ := c.Cookie("Jwt-User")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	claims := &Claims{}
	token, er := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil
	})

	if er != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Set("id", claims.UserID)
	c.Next()
}
