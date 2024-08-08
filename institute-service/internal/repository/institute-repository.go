package repository

import (
	"errors"

	"github.com/rishad004/learning-platform-go/institute-service/internal/model"
	"gorm.io/gorm"
)

type InstituteRepo interface {
	FindByAdminname(admin model.Amdin) (model.Amdin, error)
	CreateCourse(course model.Courses) error
	FetchCourses() ([]model.Courses, error)
}

type instituteRepo struct {
	DB *gorm.DB
}

func NewInstituteRepository(db *gorm.DB) InstituteRepo {
	return &instituteRepo{DB: db}
}

func (r *instituteRepo) FindByAdminname(admin model.Amdin) (model.Amdin, error) {
	var check model.Amdin
	if err := r.DB.First(&check, "Username=?", admin.Username).Error; err != nil {
		return model.Amdin{}, err
	}
	if admin.Password != check.Password {
		return model.Amdin{}, errors.New("password mismatch")
	}
	return check, nil
}

func (r *instituteRepo) CreateCourse(course model.Courses) error {
	if err := r.DB.Create(&course).Error; err != nil {
		return err
	}
	return nil
}

func (r *instituteRepo) FetchCourses() ([]model.Courses, error) {
	var courses []model.Courses
	if err := r.DB.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}
