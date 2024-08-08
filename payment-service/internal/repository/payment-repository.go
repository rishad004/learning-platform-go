package repository

import "gorm.io/gorm"

type PaymentRepo interface {
}

type paymentRepo struct {
	DB *gorm.DB
}

func NewPaymentRepository(Db *gorm.DB) PaymentRepo {
	return &paymentRepo{DB: Db}
}
