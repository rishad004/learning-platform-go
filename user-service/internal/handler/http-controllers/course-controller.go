package httpcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishad004/learning-platform-go/user-service/pkg"
)

func (u *UserHandler) CourseRouters(r *gin.Engine) {
	r.GET("/user/course", pkg.Middleware, u.GetCourseInfo)
}

func (u *UserHandler) GetCourseInfo(c *gin.Context) {
	courses, err := u.Service.GetCourseInfo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't fetch courses!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "course": courses})
}
