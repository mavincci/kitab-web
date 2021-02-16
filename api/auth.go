package api

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mavincci/Kitab-web/api/model"
	"github.com/mavincci/Kitab-web/db"
)

func UserForgot(ctx *gin.Context) {
	var user model.User

	if uname, ok := ctx.GetPostForm("uname"); ok {
		user.UserName = strings.ToLower(uname)
	} else if pno, ok := ctx.GetPostForm("pno"); ok {
		user.Pno = pno
	} else if email, ok := ctx.GetPostForm("email"); ok {
		user.Email = email
	} else {
		jsonNotFound(ctx, "id")
		return
	}
}

func UserLogin(ctx *gin.Context) {
	var user model.User

	//role, ok := ctx.GetPostForm("role")
	//if !ok {
	//	jsonNotFound(ctx, "role")
	//	return
	//}

	//id, ok := ctx.GetPostForm("uname")
	if uname, ok := ctx.GetPostForm("uname"); ok {
		user.UserName = strings.ToLower(uname)
	} else if pno, ok := ctx.GetPostForm("pno"); ok {
		user.Pno = pno
	} else if email, ok := ctx.GetPostForm("email"); ok {
		user.Email = email
	} else {
		jsonNotFound(ctx, "id")
		return
	}

	passwd, ok := ctx.GetPostForm("passwd")
	if !ok {
		jsonNotFound(ctx, "passwd")
		return
	}
	user.PasswordDigest = fmt.Sprintf("%x", md5.Sum([]byte(passwd)))

	var temp model.User
	db.DB.Where(user).First(&temp)

	if temp.PasswordDigest == "" {
		jsonNotFound(ctx, "user")
		return
	}

	//if role != temp.Role {
	//	ctx.AbortWithStatusJSON(
	//		http.StatusBadRequest,
	//		gin.H {
	//			"message": "Role doesn't MATCH!!!",
	//		})
	//	return
	//}

	auth := model.Auth{
		Token:  fmt.Sprintf("%x", md5.Sum([]byte(temp.PasswordDigest+time.Now().String()))),
		UserID: temp.ID,
		User:   temp,
	}

	db.DB.Create(&auth)

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"token": auth.Token,
			//"id": auth.UserID,
			"role":  temp.Role,
			"uname": temp.UserName,
			"email": temp.Email,
		})
}

func UserLogout(ctx *gin.Context) {
	token, ok := ctx.GetPostForm("token")
	if ok {
		auth := &model.Auth{}
		db.DB.Where(&model.Auth{Token: token}).First(&auth)
		if auth.UserID <= 0 {
			jsonUnAuthorized(ctx)
			return
		}
		db.DB.Delete(&auth)
		return
	}
	jsonNotFound(ctx, "Access Token")
}

func userAuthorize(ctx *gin.Context) (id uint, ok bool) {
	token, ok := ctx.GetPostForm("token")
	if ok {
		auth := &model.Auth{}
		db.DB.Where(&model.Auth{Token: token}).First(&auth)
		if auth.UserID <= 0 {
			jsonUnAuthorized(ctx)
			return 0, false
		}
		return auth.UserID, true
	}
	jsonNotFound(ctx, "Access Token")
	return 0, false
}
