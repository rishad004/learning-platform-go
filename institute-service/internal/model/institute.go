package model

import "gorm.io/gorm"

type Amdin struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Courses struct {
	gorm.Model
	Name  string `json:"course"`
	Price int    `json:"price"`
}
