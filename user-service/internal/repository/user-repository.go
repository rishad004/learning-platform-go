package repository

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rishad004/learning-platform-go/user-service/internal/model"
	"github.com/rishad004/learning-platform-go/user-service/pkg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user model.Users) error
	FindByUsername(user model.Users) error
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

func (r *userRepo) FindByUsername(user model.Users) error {
	var c *gin.Context
	var check model.Users
	if err := r.DB.First(&check, "Username=?", user.Username).Error; err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(check.Password), []byte(user.Password)); err != nil {
		return err
	}
	token, err := pkg.JwtCreate(check.ID, check.Email)
	if err != nil {
		return err
	}
	c.SetCookie("Jwt-User", token, int((time.Hour * 1).Seconds()), "/", "localhost", false, false)
	return nil
}
