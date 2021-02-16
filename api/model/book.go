package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model

	AuthorID		uint	`json:"-"`
	User		User `gorm:"foreignkey:AuthorID" json:"-"`

	Title		string	`json:"title"`
	Price		float64	`gorm:"default:0" json:"price"`

	Download	int32	`gorm:"default:0" json:"download"`

	Rate		int32	`gorm:"default:0" json:"rate"`	// No of rates
	AvgRating	float32	`gorm:"default:0" json:"avg_rating"`

	//Thumbnail	string
	Description	string	`gorm:"default:0" json:"description"`
	//Location	string
}
