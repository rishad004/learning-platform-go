package httpcontrollers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rishad004/learning-platform-go/enrollment-service/internal/service"
	"github.com/rishad004/learning-platform-go/enrollment-service/pkg"
)

type EnrollmentHandler struct {
	Service service.EnrollmentService
}

func NewEnrollemntController(svc service.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{Service: svc}
}

func (h *EnrollmentHandler) EnrollmentRouters(r *gin.Engine) {
	r.POST("/booking", pkg.Middleware, h.BookCourse)
}

func (h *EnrollmentHandler) BookCourse(c *gin.Context) {
	Id := c.Query("id")

	id, err := strconv.Atoi(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body!"})
		return
	}
	transaction_id, er := h.Service.BookCourse(id)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't enroll the course!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"razor_id": transaction_id})
}
