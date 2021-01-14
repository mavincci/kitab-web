package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mavincci/Kitab-web/api/model"
	"github.com/mavincci/Kitab-web/db"
	"net/http"
)

func UserUpdate(ctx *gin.Context) {
	id, ok := userAuthorize(ctx)
	if !ok {
		return
	}
	temp := &model.User{}

	db.DB.Raw("SELECT * FROM users WHERE id = ? ", id).Scan(temp)

	ctx.JSON(
		http.StatusOK,
		gin.H {
			"user": temp.UserName,
			"message": "fine",
		})
}
