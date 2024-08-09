package grpchandlers

import (
	"github.com/rishad004/learning-platform-go/enrollment-service/internal/service"
	pb "github.com/rishad004/learning-platform-go/enrollment-service/proto"
)

type EnrollmentHandler struct {
	pb.UnimplementedEnrollmentServiceServer
	Service service.EnrollmentService
}

func NewEnrollemntHandler(svc service.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{Service: svc}
}
