package service

import (
	"context"

	"github.com/rishad004/learning-platform-go/payment-service/internal/model"
	institute_pb "github.com/rishad004/learning-platform-go/payment-service/proto/client/institute"
)

type PaymentRepo interface {
	RazorIdCreate(courses *institute_pb.CourseInfoResponse, course int, userId int) (string, error)
	PaymentVerification(razor model.Razor) error
}

type PaymentService interface {
	ProcessPayment(course int, userId int) (string, error)
	RazorVerify(razor model.Razor) error
}

type paymentService struct {
	instituteClient institute_pb.AdminServiceClient
	repo            PaymentRepo
}

func NewPaymentService(repo PaymentRepo, institute institute_pb.AdminServiceClient) *paymentService {
	return &paymentService{repo: repo, instituteClient: institute}
}

func (s *paymentService) ProcessPayment(course int, userId int) (string, error) {
	courses, err := s.instituteClient.GetCourseInfo(context.Background(), &institute_pb.GetCourseInfoRequest{})
	if err != nil {
		return "", err
	}
	razorKey, err := s.repo.RazorIdCreate(courses, course, userId)
	if err != nil {
		return "", err
	}
	return razorKey, nil
}

func (s *paymentService) RazorVerify(razor model.Razor) error {
	if err := s.repo.PaymentVerification(razor); err != nil {
		return err
	}
	return nil
}
