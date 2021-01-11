package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mavincci/Kitab-web/api"
)

func main() {
	server := gin.Default()
	//server.Static("/assets/", "./assets/")

	//server.Use(middleware.ApiResolution)

	for route, handler := range api.Routes {
		server.POST(route, handler)
	}

	//for route, handler := range view.Routes {
	//	server.GET(route, handler)
	//}

	_ = server.Run(":80")
}
