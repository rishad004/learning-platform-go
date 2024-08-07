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
	u.Service.Signup(user)
}
