package router

import "github.com/gin-gonic/gin"

func  Register(app *gin.Engine) *gin.Engine {

	user := app.Group("/api/v1/user")
	{
		user.GET("/reg", func(c *gin.Context) {
			c.JSON(200,gin.H{
				"msg":"我很好",
			})
		})
	}



	return app
}
