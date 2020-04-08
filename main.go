package main

import (
	"blog-server/router"
	"github.com/gin-gonic/gin"
)

func main()  {
	app := gin.New()

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	router.Register(app)

	app.Run(":9999" )
	
}