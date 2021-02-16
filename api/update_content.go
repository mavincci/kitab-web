package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mavincci/Kitab-web/api/model"
	"github.com/mavincci/Kitab-web/db"
	"strconv"
)


func ContentUpdate(ctx *gin.Context) {
	id, ok := userAuthorize(ctx)
	if !ok {
		return
	}

	bkid, ok := ctx.GetPostForm("id")
	if !ok {
		jsonNotFound(ctx, "bkid")
		return
	}

	var curr model.User
	db.DB.Where("id = ?", id).First(&curr)

	if curr.ID == 0 {
		fmt.Println("H****************")
		jsonNotFound(ctx, "user")
		return
	}

	var currbk model.Book
	db.DB.Where("id = ?", bkid).First(&currbk)

	if currbk.ID == 0 {
		fmt.Println("H****************")
		jsonNotFound(ctx, "book")
		return
	}

	title, ok := ctx.GetPostForm("title")
	if ok {
		currbk.Title = title
	}

	desc, ok := ctx.GetPostForm("desc")
	if ok {
		currbk.Description = desc
	}

	priceStr, ok := ctx.GetPostForm("price")
	var price float64 = 0
	var err error
	if !ok {
		price = 0
	} else {
		if price, err = strconv.ParseFloat(priceStr, 64); err != nil {
			price = 0
		}
	}

	if price != 0 {
		currbk.Price = price
	} else {
		currbk.Price = 0
	}

	db.DB.Save(&currbk)
}



func ContentDelete(ctx *gin.Context) {
	id, ok := userAuthorize(ctx)
	if !ok {
		return
	}

	bkid, ok := ctx.GetPostForm("id")
	if !ok {
		jsonNotFound(ctx, "bkid")
		return
	}

	var curr model.User
	db.DB.Where("id = ?", id).First(&curr)

	if curr.ID == 0 {
		fmt.Println("H****************")
		jsonNotFound(ctx, "user")
		return
	}

	var currbk model.Book
	db.DB.Where("id = ?", bkid).First(&currbk)

	if currbk.ID == 0 {
		fmt.Println("H****************")
		jsonNotFound(ctx, "book")
		return
	}

	fmt.Print("ID : ")
	fmt.Println(currbk.AuthorID)
	if curr.ID == currbk.AuthorID || curr.Role == "admin" {
		db.DB.Delete(&currbk)
		return
	}
	jsonUnAuthorized(ctx)
}
