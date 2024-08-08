package httpcontrollers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rishad004/learning-platform-go/institute-service/internal/model"
	"github.com/rishad004/learning-platform-go/institute-service/internal/service"
	"github.com/rishad004/learning-platform-go/institute-service/pkg"
)

type InstituteHandler struct {
	Service service.InstituteService
}

func (h *InstituteHandler) AdminRouters(r *gin.Engine) {
	r.POST("/login/admin", h.Login)
	r.POST("/course", h.AddCourse)
	r.GET("/course", h.GetCourseInfo)
}

func NewInstituteController(svc service.InstituteService) *InstituteHandler {
	return &InstituteHandler{Service: svc}
}

func (h *InstituteHandler) Login(c *gin.Context) {
	var admin model.Amdin

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body!"})
		return
	}
	check, err := h.Service.Login(admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password!"})
		return
	}
	token, err := pkg.JwtCreate(check.ID, check.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't login try again later"})
	}

	c.SetCookie("Jwt-Admin", token, int((time.Hour * 1).Seconds()), "/", "localhost", false, false)

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

func (h *InstituteHandler) GetCourseInfo(c *gin.Context) {
	courses, err := h.Service.GetCourseInfo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't fetch the course data!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"courses": courses})
}
