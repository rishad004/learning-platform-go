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
