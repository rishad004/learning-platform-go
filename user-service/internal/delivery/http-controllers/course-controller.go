package httpcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserHandler) GetCourseInfo(c *gin.Context) {
	courses, err := u.Service.GetCourseInfo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't fetch courses!"})
		return
	}
	c.JSON(http.StatusOK, courses)
}
