package model

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	UserId        uint
	OrderId       string
	PaymentId     string
	CourseId      uint
	PaymentStatus bool
}

type Razor struct {
	Order     string `json:"OrderID"`
	Payment   string `json:"PaymentID"`
	Signature string `json:"Signature"`
}