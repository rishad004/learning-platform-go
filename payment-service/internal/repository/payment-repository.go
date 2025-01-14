package repository

import (
	"errors"

	"github.com/rishad004/learning-platform-go/payment-service/internal/model"
	"github.com/rishad004/learning-platform-go/payment-service/pkg"
	institute_pb "github.com/rishad004/learning-platform-go/payment-service/proto/client/institute"
	"gorm.io/gorm"
)

type PaymentRepo interface {
	RazorIdCreate(courses *institute_pb.CourseInfoResponse, course int, userId int) (string, error)
	PaymentVerification(razor model.Razor) error
}

type paymentRepo struct {
	DB *gorm.DB
}

func NewPaymentRepository(Db *gorm.DB) PaymentRepo {
	return &paymentRepo{DB: Db}
}

func (r *paymentRepo) RazorIdCreate(courses *institute_pb.CourseInfoResponse, course int, userId int) (string, error) {
	for _, v := range courses.Courses {
		if v.Id == int32(course) {
			razorKey, err := pkg.Executerazorpay(v.Course, int(v.Price))
			if err != nil {
				return "", err
			}
			r.DB.Create(&model.Booking{
				UserId:        uint(userId),
				OrderId:       razorKey,
				CourseId:      uint(v.Id),
				PaymentStatus: false,
			})
			return razorKey, nil
		}
	}
	return "", errors.New("course not found")
}

func (r *paymentRepo) PaymentVerification(razor model.Razor) error {
	if err := pkg.RazorPaymentVerification(razor.Signature, razor.Order, razor.Payment); err != nil {
		return err
	}
	if err := r.DB.Model(&model.Booking{}).Where("order_id = ?", razor.Order).Update("payment_id", razor.Payment).Error; err != nil {
		return err
	}
	if err := r.DB.Model(&model.Booking{}).Where("order_id = ?", razor.Order).Update("payment_status", true).Error; err != nil {
		return err
	}
	return nil
}
