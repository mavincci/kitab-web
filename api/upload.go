package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mavincci/Kitab-web/api/model"
	"github.com/mavincci/Kitab-web/db"
	"strconv"
	"strings"
)


func ContentUpload(ctx *gin.Context) {
	tmbnl, err := ctx.FormFile("tmbnl")
	//tmbnl, err = ctx.FormFile("tmbnl")

	if err != nil {
		jsonNotFound(ctx, "tmbnl")
		return
	}

	cntnt, err := ctx.FormFile("cntnt")
	//file, err = ctx.FormFile("cntnt")

	if err != nil {
		jsonNotFound(ctx, "file")
		return
	}

	title, ok := ctx.GetPostForm("title")
	if !ok {
		jsonNotFound(ctx, "title")
		return
	}

	desc, ok := ctx.GetPostForm("desc")
	if !ok {
		jsonNotFound(ctx, "desc")
		return
	}

	priceStr, ok := ctx.GetPostForm("price")
	var price float64 = 0
	if !ok {
		price = 0
	} else {
		if price, err = strconv.ParseFloat(priceStr, 64); err != nil {
			price = 0
		}
	}


	id, ok := userAuthorize(ctx)
	if !ok {
		return
	}

	curr := &model.User{}
	db.DB.Where("id = ?", id).First(&curr)

	if curr.ID == 0 {
		jsonNotFound(ctx, "user")
		return
	}

	currBook := model.Book{
		AuthorID:    curr.ID,
		Title:       strings.ToLower(title),
		Price:       price,
		Download:    0,
		Rate:        0,
		AvgRating:   0,
		Description: desc,
	}

	db.DB.Create(&currBook)

	_ = ctx.SaveUploadedFile(
		tmbnl,
		ThumbLoc+ fmt.Sprintf("%d.png", currBook.ID))

	_ = ctx.SaveUploadedFile(
		cntnt,
		CntntLoc+ fmt.Sprintf("%d.pdf", currBook.ID))
}
