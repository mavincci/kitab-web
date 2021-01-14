package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	UserName		string `json:"user_name" gorm:"unique"`
	PasswordDigest	string `json:"-"`
	Email			string `json:"email" gorm:"unique"`
	Pno				string `json:"pno" gorm:"unique"`
	Role			string `json:"-"`
}
