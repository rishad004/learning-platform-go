package repository

import "gorm.io/gorm"

type EnrollemntRepo interface{}

type enrollmentRepo struct {
	Db *gorm.DB
}

func NewEnrollemntRepository(Db *gorm.DB) EnrollemntRepo {
	return &enrollmentRepo{Db: Db}
}
