package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mavincci/Kitab-web/api"
	"github.com/mavincci/Kitab-web/db"

	"github.com/gin-contrib/cors"
	//cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	server := gin.Default()
	//server.Static("/assets/", "./assets/")
	//server.Static("/assets/", "../thumb/")
	//server.Use(middleware.ApiResolution)

	server.Use(cors.Default())

	defer db.CloseDB()

	api.BuildRoutes(server)

	_ = server.Run(":80")
}
