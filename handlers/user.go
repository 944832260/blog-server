package handlers

import (
	"blog-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type UserHandlers struct {
	BaseHandler
}

type PhoneInfo struct {
	Phone string `json:"phone"`
	PassWord string `json:"password"`
	Name string `json:"name"`
	Id string `json:"id"`
} 


func (u *UserHandlers) CreateUser(c *gin.Context)  {
	var (
		phoneInfo PhoneInfo
		user models.User
	)

	_ = c.BindJSON(&phoneInfo)

	phone := phoneInfo.Phone
	password := phoneInfo.PassWord
	name := phoneInfo.Name
	fmt.Println(phone,"phone---")
	fmt.Println(password,"phone---")
	fmt.Println(name,"phone---")
	if phone == "" {
		c.JSON(http.StatusOK,gin.H{
			"code":400,
			"msg":"请填写手机号",
		})
		return
	}
	if password == ""{
		c.JSON(http.StatusOK,gin.H{
			"code":400,
			"msg":"请填写密码",
		})
		return
	}
	//连接数据库
	db := models.GetDB()
	//查询有没有这个手机号
	db.Where(models.User{Tel:phone}).First(&user)
	if user == (models.User{}) {
		user = models.User{Tel:phone,PassWord:password,Name:name}
		db.Create(&user)
	} else {
		c.JSON(http.StatusOK,gin.H{
			"code":400,
			"msg":"用户已存在",
		})

		db.Close()
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"添加成功",
	})
}

func (u *UserHandlers) DeleteUser(c *gin.Context)  {
	var (
		phoneInfo PhoneInfo
		user models.User
	)

	_ = c.Bind(&phoneInfo)
	phone := phoneInfo.Phone

	fmt.Println(phone,"1111111")
	if phone == "" {
		c.JSON(http.StatusOK,gin.H{
			"code":400,
			"msg":"请填写手机号",
		})
		return
	}

	//连接数据库
	db := models.GetDB()
	db.Where("tel = ?",phone).Delete(&user)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"删除成功",
	})
}