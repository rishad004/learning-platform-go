package repository

import "gorm.io/gorm"

type HotelRepo interface {
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepository(Db *gorm.DB) HotelRepo {
	return &userRepo{DB: Db}
}
