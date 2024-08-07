package httpcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishad004/learning-platform-go/institute-service/internal/model"
	"github.com/rishad004/learning-platform-go/institute-service/internal/service"
)

type InstituteHandler struct {
	Service service.InstituteService
}

func (h *InstituteHandler) AdminRouters(r *gin.Engine) {
	r.POST("/login/admin", h.Login)
	r.POST("/course", h.AddCourse)
}

func (h *InstituteHandler) Login(c *gin.Context) {
	var admin model.Amdin

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body!"})
		return
	}
	if err := h.Service.Login(admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login successful!"})
}

func (h *InstituteHandler) AddCourse(c *gin.Context) {
	var course model.Courses
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "course created successfully!"})
}
