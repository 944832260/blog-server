package router

import (
	"blog-server/handlers"
	"github.com/gin-gonic/gin"
)

func  Register(app *gin.Engine) *gin.Engine {
	UserBase := new(handlers.UserHandlers)
	user := app.Group("/api/v1/user")
	{
		user.GET("/reg", UserBase.CreateUser)
	}



	return app
}
