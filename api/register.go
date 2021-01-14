package api

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mavincci/Kitab-web/api/model"
	"github.com/mavincci/Kitab-web/db"
	"strings"
	"time"
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
		PasswordDigest: fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String() + passwd))),
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

	ctx.JSON(200, usr)
}

