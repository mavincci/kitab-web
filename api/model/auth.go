package model

import "gorm.io/gorm"

type Auth struct {
	gorm.Model		`json:"_"`
	Token	string
	UserID	uint	`json:"_"`
	User	User	`gorm:"foreignkey:UserID" json:"_"`
}
