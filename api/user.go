package api

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mavincci/Kitab-web/api/model"
	"github.com/mavincci/Kitab-web/db"
)

func UserRegister(context *gin.Context) {
	uname := context.PostForm("uname")
	pno := context.PostForm("pno")
	email := context.PostForm("email")
	passwd := context.PostForm("passwd")

	usr := model.User{
		UserName:       uname,
		PasswordDigest: fmt.Sprintf("%x", md5.Sum([]byte(passwd))),
		Email:          email,
		Pno:            pno,
	}

	db.DB.NewRecord(&usr)
	db.DB.Create(&usr)

	context.JSON(200, usr)
}

func UserLogin(context *gin.Context) {
	//user_name, _ := context.GetPostForm("uname")

}
