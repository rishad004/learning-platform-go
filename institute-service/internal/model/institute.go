package model

import "gorm.io/gorm"

type Amdin struct {
	gorm.Model
	Username string
	Password string
}
