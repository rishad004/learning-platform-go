package service

import (
	"context"

	institute_pb "github.com/rishad004/learning-platform-go/payment-service/proto/client/institute"
)

type PaymentRepo interface {
	RazorIdCreate(courses *institute_pb.CourseInfoResponse, course int) (string, error)
}

type PaymentService interface {
	ProcessPayment(course int) (string, error)
}

type paymentService struct {
	instituteClient institute_pb.AdminServiceClient
	repo            PaymentRepo
}

func NewPaymentService(repo PaymentRepo, institute institute_pb.AdminServiceClient) *paymentService {
	return &paymentService{repo: repo, instituteClient: institute}
}

func (s *paymentService) ProcessPayment(course int) (string, error) {
	courses, err := s.instituteClient.GetCourseInfo(context.Background(), &institute_pb.GetCourseInfoRequest{})
	if err != nil {
		return "", err
	}
	razorKey, err := s.repo.RazorIdCreate(courses, course)
	if err != nil {
		return "", err
	}
	return razorKey, nil
}
