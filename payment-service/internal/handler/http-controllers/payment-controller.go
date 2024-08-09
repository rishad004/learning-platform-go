package httpcontrollers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rishad004/learning-platform-go/payment-service/internal/model"
	"github.com/rishad004/learning-platform-go/payment-service/internal/service"
)

type PaymentHandler struct {
	Service service.PaymentService
}

func (h *PaymentHandler) PaymentRouters(r *gin.Engine) {
	r.GET("/payment", h.RazorRender)
	r.POST("/payment", h.RazorVerify)
}

func NewPaymentController(svc service.PaymentService) *PaymentHandler {
	return &PaymentHandler{Service: svc}
}

func (h *PaymentHandler) RazorRender(c *gin.Context) {
	orderId := c.Query("id")

	c.HTML(200, "razor.html", gin.H{
		"Order":   orderId,
		"Key":     os.Getenv("RAZOR_KEY"),
		"Product": "Your course is ready. Pay for them now!",
	})
}

func (h *PaymentHandler) RazorVerify(c *gin.Context) {
	var razor model.Razor

	if err := c.ShouldBindJSON(&razor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := h.Service.RazorVerify(razor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id expired or invalid!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "payment successful!"})
}
