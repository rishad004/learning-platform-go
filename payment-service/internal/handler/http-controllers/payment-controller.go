package httpcontrollers

import (
	"github.com/gin-gonic/gin"
	"github.com/rishad004/learning-platform-go/payment-service/internal/service"
)

type PaymentHandler struct {
	Service service.PaymentService
}

func (h *PaymentHandler) PaymentRouters(r *gin.Engine) {

}

func NewPaymentController(svc service.PaymentService) *PaymentHandler {
	return &PaymentHandler{Service: svc}
}
