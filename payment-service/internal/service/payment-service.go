package service

type PaymentRepo interface{}

type PaymentService interface{}

type paymentService struct {
	repo PaymentRepo
}

func NewPaymentService(repo PaymentRepo) *paymentService {
	return &paymentService{repo: repo}
}
