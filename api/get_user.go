package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mavincci/Kitab-web/api/model"
	"github.com/mavincci/Kitab-web/db"
)

func GetAuthors(ctx *gin.Context) {
	id, ok := userAuthorize(ctx)
	if !ok {
		return
	}

	curr := &model.User{}
	db.DB.Where("id = ?", id).First(&curr)

	if curr.ID == 0 {
		fmt.Println("H****************")
		jsonNotFound(ctx, "user")
		return
	}

	fmt.Println(curr.Role)
	if curr.Role != "admin" {
		jsonUnAuthorized(ctx)
		return
	}

	var authors []model.User
	db.DB.Where("role = ?", "author").Find(&authors)

	ctx.JSON(200, authors)
}

func GetPublishers(ctx *gin.Context) {
	id, ok := userAuthorize(ctx)
	if !ok {
		return
	}

	curr := &model.User{}
	db.DB.Where("id = ?", id).First(&curr)

	if curr.ID == 0 {
		fmt.Println("H****************")
		jsonNotFound(ctx, "user")
		return
	}

	fmt.Println(curr.Role)
	if curr.Role != "admin" {
		jsonUnAuthorized(ctx)
		return
	}

	var publishers []model.User
	db.DB.Where("role = ?", "publisher").Find(&publishers)

	ctx.JSON(200, publishers)
}
