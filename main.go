package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default(); //返回的路由引擎
	//指定用户使用某种方式请求，接收两个参数，1.路由2.执行函数
	router.GET("/hello",sayHello)
	//启动服务
	router.Run(":9090")
}