package grpchandlers

import (
	"github.com/rishad004/learning-platform-go/payment-service/internal/service"
	pb "github.com/rishad004/learning-platform-go/payment-service/proto"
)

type PaymentHandler struct {
	pb.UnimplementedPaymentServiceServer
	Service service.PaymentService
}

func NewPaymentHandler(svc service.PaymentService) *PaymentHandler {
	return &PaymentHandler{Service: svc}
}
