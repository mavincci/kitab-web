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
	engine.POST("/api/logout", UserLogout)
	engine.POST("/api/register", UserRegister)
	engine.POST("/api/delete", UserDelete)


	engine.POST("/api/user/update", UserUpdate)
	engine.POST("/api/user/get", GetUser)

	engine.POST("/api/user/getauthors", GetAuthors)
	engine.POST("/api/user/getpublishers", GetPublishers)

	//engine.POST("/api/content/get", GetBook)
	engine.POST("/api/content/upload", ContentUpload)
	engine.POST("/api/content/search", ContentSearch)
	engine.POST("/api/content/searchbyme", ContentSearchByAuthor)
	engine.POST("/api/content/byme", ContentByAuthor)
	engine.POST("/api/content/bymetoprated", ContentByAuthorTopRated)
	engine.POST("/api/content/bymetopdown", ContentByAuthorTopDown)
	engine.POST("/api/content/update", ContentUpdate)
	engine.POST("/api/content/delete", ContentDelete)
	engine.POST("/api/content/random", GetRandom)


	engine.GET("/", func(c *gin.Context) {
		c.String(200,
			"This is kitab .com")
	})
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
	fmt.Println("empty ", col)
	ctx.JSON(
		400,
		gin.H {
			//"time-stamp": time.Now().Format(time.RFC822),
			"message": col + " not found",
			"error": "not found",
			"attribute": col,
		})
}

func jsonHeld(ctx *gin.Context, col string) {
	fmt.Println("already held ", col)
	ctx.JSON(
		400,
		gin.H {
			//"time-stamp": time.Now().Format(time.RFC822),
			"message": col + " already held",
			"error": "already held",
			"attribute": col,
		})
}
