package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {

	ctx.Status(http.StatusOK)
	ctx.Writer.Write([]byte("king"))

	//firstName, _ := ctx.GetPostForm("firstName")
	//middleName, _ := ctx.GetPostForm("middleName")
	//lastName, _ := ctx.GetPostForm("lastName")
	//age, _ := ctx.GetPostForm("age")
	//nationality, _ := ctx.GetPostForm("nationality")
	//eyeColor, _ := ctx.GetPostForm("eyeColor")
	//
	//ctx.JSON(
	//	200,
	//	gin.H{
	//		"firstName":   firstName,
	//		"middleName":  middleName,
	//		"lastName":    lastName,
	//		"age":         age,
	//		"nationality": nationality,
	//		"eyeColor":    eyeColor,
	//		"message":     "Success full try",
	//	})
}
