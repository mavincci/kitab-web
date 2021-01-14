package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var DB *gorm.DB
var dbErr error

func init() {
	buildDb()
}

func buildDb() {
	host	:= os.Getenv("db_host")
	port	:= os.Getenv("db_port")
	user	:= os.Getenv("db_user")
	dbname	:= os.Getenv("db_name")
	pass	:= os.Getenv("db_pass")
	ssl		:= os.Getenv("db_ssl")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s %s",
		host, port, user, dbname, pass, ssl)

	DB, dbErr = gorm.Open("postgres", dsn)
	if dbErr != nil {
		fmt.Println("Can not connect to db")
		fmt.Println(dbErr.Error())
		os.Exit(0)
	} else {
		fmt.Println("Database successfully connected.")
	}
}

func CloseDB()  {
	err := DB.Close()
	if err != nil {
		panic(err)
	}
}
