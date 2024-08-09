package grpchandlers

import (
	"context"

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

func (h *PaymentHandler) ProcessPayment(c context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
	razorKey, err := h.Service.ProcessPayment(int(req.Course), int(req.Price))
	if err != nil {
		return nil, err
	}
	return &pb.ProcessPaymentResponse{TransactionId: razorKey}, nil
}
