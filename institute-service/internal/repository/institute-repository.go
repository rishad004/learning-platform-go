package repository

import (
	"errors"

	"github.com/rishad004/learning-platform-go/institute-service/internal/model"
	"gorm.io/gorm"
)

type InstituteRepo interface {
	FindByAdminname(admin model.Amdin) error
}

type instituteRepo struct {
	DB *gorm.DB
}

func NewInstituteRepository(db *gorm.DB) InstituteRepo {
	return &instituteRepo{DB: db}
}

func (r *instituteRepo) FindByAdminname(admin model.Amdin) error {
	var check model.Amdin
	if err := r.DB.First(&check, "Username=?", admin.Username).Error; err != nil {
		return nil
	}
	if admin.Password != check.Password {
		return errors.New("password mismatch")
	}
	return nil
}
