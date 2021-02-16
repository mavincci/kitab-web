package api

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mavincci/Kitab-web/api/model"
	"github.com/mavincci/Kitab-web/db"
)


func UserRegister(ctx *gin.Context) {
	uname, ok := ctx.GetPostForm("uname")
	if !ok {
		jsonNotFound(ctx, "uname")
		return
	}
	pno, ok := ctx.GetPostForm("pno")
	if !ok {
		jsonNotFound(ctx, "pno")
		return
	}
	email, ok := ctx.GetPostForm("email")
	if !ok {
		jsonNotFound(ctx, "email")
		return
	}
	passwd, ok := ctx.GetPostForm("passwd")
	if !ok {
		jsonNotFound(ctx, "passwd")
		return
	}
	role, ok := ctx.GetPostForm("role")
	if !ok {
		jsonNotFound(ctx, "role")
		return
	}

	if role == "admin" {
		jsonHeld(ctx, "unauthorized role\"Admin\"")
		return
	}

	usr := model.User{
		UserName:       strings.ToLower(uname),
		PasswordDigest: fmt.Sprintf("%x", md5.Sum([]byte(passwd))),
		Email:          email,
		Pno:            pno,
		Role:			role,
	}

	temp := &model.User{}

	db.DB.Where("user_name = ?", usr.UserName).First(temp)
	if temp.UserName != "" {
		jsonHeld(ctx, "uname")
		return
	}
	temp.UserName = ""

	db.DB.Where("email = ?", usr.Email).First(temp)
	if temp.Email != "" {
		jsonHeld(ctx, "email")
		return
	}
	temp.Email = ""

	db.DB.Where("pno = ?", usr.Pno).First(temp)
	if temp.Pno != "" {
		jsonHeld(ctx, "pno")
		return
	}
	temp.Pno = ""

	db.DB.Create(&usr)

	//ctx.JSON(200, usr)
}


func UserDelete(ctx *gin.Context) {
	id, ok := userAuthorize(ctx)
	if !ok {
		return
	}

	curr := &model.User{}
	db.DB.Where("id = ?", id).First(&curr)

	if curr.ID == 0 {
		//fmt.Println("H****************")
		jsonNotFound(ctx, "user")
		return
	}
	var auths []model.Auth
	db.DB.Where("user_id = ?", curr.ID).Find(&auths)
	db.DB.Unscoped().Delete(auths)
	//db.DB.Delete(auths)


	db.DB.Delete(curr)
}
