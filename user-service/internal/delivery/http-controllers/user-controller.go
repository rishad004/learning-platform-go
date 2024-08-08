package httpcontrollers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rishad004/learning-platform-go/user-service/internal/model"
	"github.com/rishad004/learning-platform-go/user-service/internal/service"
	"github.com/rishad004/learning-platform-go/user-service/pkg"
)

type UserHandler struct {
	Service service.UserService
}

func (u *UserHandler) UserRouters(r *gin.Engine) {

	u.CourseRouters(r)

	r.POST("/signup", u.Signup)
	r.POST("/login", u.Login)
}

func NewUserController(svc service.UserService) *UserHandler {
	return &UserHandler{Service: svc}
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
	check, err := u.Service.Login(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials!"})
		return
	}

	token, er := pkg.JwtCreate(check.ID, check.Email)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't login try again later!!"})
		return
	}

	c.SetCookie("Jwt-User", token, int((time.Hour * 1).Seconds()), "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{"message": "successfully logged in!"})
}
