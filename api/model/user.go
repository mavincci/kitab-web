package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName       string `json:"user_name"`
	PasswordDigest string `json:"-"`
	Email          string `json:"email"`
	Pno            string `json:"pno"`
	//Role			string `json:"-"`

}
