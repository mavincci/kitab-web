package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var dbErr error

func init() {

	/*
		host := "ec2-54-247-169-129.eu-west-1.compute.amazonaws.com"
		port := 5432
		user := "xzyeyiipjjlfhh"
		dbname := "damld4k2dmoicp"
		passwd := "aebfae8e98e323ebd4d53c730ca44579c5d27a561dacea005194619a3581586c"

		dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
			host, port, user, dbname, passwd)
	*/
	host := os.Getenv("db_host")   // "localhost"
	port := os.Getenv("db_port")   // 5432
	user := os.Getenv("db_user")   //"postgres"
	dbname := os.Getenv("db_name") // "kitab_test"
	pass := os.Getenv("db_pass")   // "nil@mavpgmaster"
	ssl := os.Getenv("db_ssl")     // "sslmode=disable"

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
