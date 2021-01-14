package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
}

func BuildRoutes(engine *gin.Engine) {
	engine.POST("/api/auth", UserLogin)
	engine.POST("/api/register", UserRegister)

	engine.POST("/api/update", UserUpdate)

	engine.GET("/api/books", GetBooks)

	//engine.POST("/api/logout", Logout)
	//engine.POST("/api/update", UserUpdate)
}

func jsonUnAuthorized(ctx *gin.Context) {
	fmt.Println("Unauth tried to update")
	ctx.JSON(
		http.StatusUnauthorized,
		gin.H {
			"message": "Unauthorized user",
		})
}

func jsonNotFound(ctx *gin.Context, col string) {
	fmt.Println("not found ", col)
	ctx.JSON(
		400,
		gin.H {
			//"time-stamp": time.Now().Format(time.RFC822),
			"error": "bad request",
			"message": "not found",
			"attribute": col,
		})
}

func jsonHeld(ctx *gin.Context, col string) {
	fmt.Println("already held ", col)
	ctx.JSON(
		400,
		gin.H {
			//"time-stamp": time.Now().Format(time.RFC822),
			"error": "bad request",
			"message": "already held",
			"attribute": col,
		})
}
