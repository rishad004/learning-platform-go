package service

import (
	"context"

	payment_pb "github.com/rishad004/learning-platform-go/enrollment-service/proto/client/payment"
)

type EnrollemntRepo interface{}

type EnrollmentService interface {
	BookCourse(course int) (string, error)
}

type enrollmentService struct {
	paymentClient payment_pb.PaymentServiceClient
	repo          EnrollemntRepo
}

func NewEnrollemntService(repo EnrollemntRepo, payment payment_pb.PaymentServiceClient) EnrollmentService {
	return &enrollmentService{repo: repo, paymentClient: payment}
}

func (s *enrollmentService) BookCourse(course int) (string, error) {
	transaction_id, err := s.paymentClient.ProcessPayment(context.Background(), &payment_pb.ProcessPaymentRequest{Course: int32(course)})
	if err != nil {
		return "", err
	}
	return transaction_id.TransactionId, nil
}
