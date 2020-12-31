package main

import (
	"github.com/gin-gonic/gin"
	"siteend/handlers"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handlers.Index)
	router.GET("/:path", handlers.Index)
	router.GET("/:path/:path1", handlers.Index)
	router.GET("/:path/:path1/:path2", handlers.Index)
	router.GET("/:path/:path1/:path2/:path3", handlers.Index)
	//router.GET("/ping", handlers.Ping)
	router.Run(":8089")
	// listen and serve on 0.0.0.0:8089
}

