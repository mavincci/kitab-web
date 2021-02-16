package model

import "gorm.io/gorm"

type Trade struct {
	gorm.Model		`json:"-"`
	BookId	uint
	Book	Book	`gorm:"foreignkey:BookID" json:"_"`
	UserId	uint
	User	User	`gorm:"foreignkey:UserID" json:"_"`
	Rate	float32
	Comment	string
}
