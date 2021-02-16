package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mavincci/Kitab-web/api/model"
	"github.com/mavincci/Kitab-web/db"
	"net/http"
	"strings"
)

func ContentByAuthor(ctx *gin.Context) {
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

	fmt.Println(curr.Role)

	if curr.Role == "publisher" || curr.Role == "author" {
		var books []model.Book
		db.DB.Where("author_id = ?", curr.ID).Find(&books)

		ctx.JSON(
			http.StatusOK,
			books)
		return
	}
	jsonUnAuthorized(ctx)
}

func ContentSearchByAuthor(ctx *gin.Context) {
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

	if curr.Role != "publisher" && curr.Role != "author" {
		jsonUnAuthorized(ctx)
		return
	}

	title, ok := ctx.GetPostForm("title")
	if !ok {
		jsonNotFound(ctx, "bktitle")
		return
	}

	title = strings.ToLower(title)

	var books []model.Book
	db.DB.Where("title LIKE ?", "%"+title+"%").Find(&books)

	//
	//ctx.JSON(
	//	http.StatusOK,
	//	books)

	type res struct {
		ID uint			`json:"id"`
		Title		string	`json:"title"`
		Price		float64	`gorm:"default:0" json:"price"`

		Download	int32	`gorm:"default:0" json:"download"`

		Rate		int32	`gorm:"default:0" json:"rate"`	// No of rates
		AvgRating	float32	`gorm:"default:0" json:"avg_rating"`

		//Thumbnail	string
		Description	string	`gorm:"default:0" json:"description"`
		AuthorName string 	`json:"author"`
	}

	var t []res
	rows , _:= db.DB.Table("books").
		Select(" books.id, books.title, books.price, books.download, books.rate, books.avg_rating, books.description, users.user_name").
		Joins("left join users on users.id = books.author_id").Rows()



	for rows.Next() {
		var ID uint
		var Title		string
		var Price		float64
		var Download	int32
		var Rate		int32
		var AvgRating	float32
		var Description	string
		var Author string

		rows.Scan(&ID, &Title, &Price, &Download, &Rate, &AvgRating, &Description, &Author)
		if strings.Contains(Title, title) {
			t = append(t, res{
				ID, Title, Price, Download, Rate, AvgRating, Description, Author,
			})
		}
	}

	ctx.JSON(
		http.StatusOK,
		t)
}

func ContentSearch(ctx *gin.Context) {
	title, ok := ctx.GetPostForm("title")
	if !ok {
		jsonNotFound(ctx, "bktitle")
		return
	}

	title = strings.ToLower(title)

	//var books []model.Book
	//db.DB.Where("title LIKE ?", "%"+title+"%").Find(&books)

	//
	//ctx.JSON(
	//	http.StatusOK,
	//	books)

	type res struct {
		ID uint			`json:"id"`
		Title		string	`json:"title"`
		Price		float64	`gorm:"default:0" json:"price"`

		Download	int32	`gorm:"default:0" json:"download"`

		Rate		int32	`gorm:"default:0" json:"rate"`	// No of rates
		AvgRating	float32	`gorm:"default:0" json:"avg_rating"`

		//Thumbnail	string
		Description	string	`gorm:"default:0" json:"description"`
		AuthorName string 	`json:"author"`
	}

	var t []res
	rows , _:= db.DB.Table("books").
		Select(" books.id, books.title, books.price, books.download, books.rate, books.avg_rating, books.description, users.user_name").
		Joins("left join users on users.id = books.author_id").Rows()

	for rows.Next() {
		var ID uint
		var Title		string
		var Price		float64
		var Download	int32
		var Rate		int32
		var AvgRating	float32
		var Description	string
		var Author string

		rows.Scan(&ID, &Title, &Price, &Download, &Rate, &AvgRating, &Description, &Author)
		if strings.Contains(Title, title) {
			t = append(t, res{
				ID, Title, Price, Download, Rate, AvgRating, Description, Author,
			})
		}
	}

	ctx.JSON(
		http.StatusOK,
		t)
}

func ContentByAuthorTopRated(ctx *gin.Context) {
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

	if curr.Role != "publisher" && curr.Role != "author" {
		jsonUnAuthorized(ctx)
		return
	}

	var books []model.Book
	db.DB.Where("author_id = ?", curr.ID).Order("avg_rating desc").Limit(1).Find(&books)

	ctx.JSON(
		http.StatusOK,
		books)
}


func ContentByAuthorTopDown(ctx *gin.Context) {
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

	if curr.Role != "publisher" && curr.Role != "author" {
		jsonUnAuthorized(ctx)
		return
	}

	var books []model.Book
	db.DB.Where("author_id = ?", curr.ID).Order("download desc").Limit(1).Find(&books)

	ctx.JSON(
		http.StatusOK,
		books)

}

func GetRandom(ctx *gin.Context) {
//	var books []model.Book
//	db.DB.Not("author_id = ?", 0).Find(&books)
}

