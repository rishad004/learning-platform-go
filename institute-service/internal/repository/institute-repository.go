package repository

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rishad004/learning-platform-go/institute-service/internal/model"
	"github.com/rishad004/learning-platform-go/institute-service/pkg"
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
	var c *gin.Context
	if err := r.DB.First(&check, "Username=?", admin.Username).Error; err != nil {
		return nil
	}
	if admin.Password != check.Password {
		return errors.New("password mismatch")
	}
	token, err := pkg.JwtCreate(admin.ID, admin.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't login try again later"})
	}

	c.SetCookie("Jwt-Admin", token, int((time.Hour * 1).Seconds()), "/", "localhost", false, false)
	return nil
}
