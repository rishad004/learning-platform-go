package repository

import (
	"github.com/rishad004/learning-platform-go/user-service/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user model.Users) error
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepository(Db *gorm.DB) UserRepo {
	return &userRepo{DB: Db}
}

func (r *userRepo) CreateUser(user model.Users) error {
	hashedpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}
	
	user.Password = string(hashedpass)

	err = r.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
