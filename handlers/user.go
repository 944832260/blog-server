package handlers

import (
	"blog-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandlers struct {
	BaseHandler
}



func (u *UserHandlers) CreateUser(c *gin.Context)  {
	models.GetDB()
	c.JSON(http.StatusOK,gin.H{
		"msg":"0006",
	})
}