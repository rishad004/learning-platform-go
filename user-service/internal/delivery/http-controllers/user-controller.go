package httpcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishad004/learning-platform-go/user-service/internal/model"
	"github.com/rishad004/learning-platform-go/user-service/internal/service"
)

type UserHandler struct {
	Service service.UserService
}

func (u *UserHandler) UserRouters(r *gin.Engine) {
	r.POST("/signup", u.Signup)
}

func (u *UserHandler) Signup(c *gin.Context) {
	var user model.Users

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body!"})
		return
	}
	if err := u.Service.Signup(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't create user!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created successfully!"})
}

func (u *UserHandler) Login(c *gin.Context) {
	var user model.Users

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged in!"})
}
