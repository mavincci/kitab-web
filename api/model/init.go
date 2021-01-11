package model

import "github.com/mavincci/Kitab-web/db"

func init() {
	db.DB.AutoMigrate(&User{})
}
