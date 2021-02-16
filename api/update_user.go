package api

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mavincci/Kitab-web/api/model"
	"github.com/mavincci/Kitab-web/db"
	"net/http"
)


func GetUser(ctx *gin.Context) {
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

	ctx.JSON(
		http.StatusOK,
		curr)
}

func UserUpdate(ctx *gin.Context) {
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

	/*	fmt.Println(curr.UserName)
		fmt.Println(curr.PasswordDigest)
		fmt.Println(curr.Email)
		fmt.Println(curr.Pno)
		fmt.Println(curr.Role)
	return*/

	temp := &model.User{}

	uname, ok := ctx.GetPostForm("uname")
	if ok {
		if uname != curr.UserName {
			db.DB.Where("user_name = ?", uname).First(temp)
			if temp.UserName != "" {
				jsonHeld(ctx, "uname")
				return
			}
			curr.UserName = uname
		}
	}

	pno, ok := ctx.GetPostForm("pno")
	if ok {
		if pno != curr.Pno {
			db.DB.Where("pno = ?", pno).First(temp)
			if temp.Pno != "" {
				jsonHeld(ctx, "pno")
				return
			}
			curr.Pno = pno
		}
	}

	email, ok := ctx.GetPostForm("email")
	if ok {
		if email != curr.Email {
			db.DB.Where("email = ?", email).First(temp)
			if temp.Email != "" {
				jsonHeld(ctx, "email")
				return
			}
			curr.Email = email
		}
	}

	passwd, ok := ctx.GetPostForm("passwd")
	if ok {
		digest := fmt.Sprintf("%x", md5.Sum([]byte(passwd)))
		if digest != curr.PasswordDigest {
			curr.PasswordDigest = digest
		}
	}

	db.DB.Save(&curr)


	//ctx.JSON(
	//	http.StatusOK,
	//	gin.H {
	//		"user": temp.UserName,
	//		"message": "fine",
	//	})
}


