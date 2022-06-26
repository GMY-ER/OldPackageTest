package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//使用默认中间件创建一个gin路由器
	//logger and recovery (crash-free)中间件
	router := gin.Default()
	//restful的开发中
	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", puting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", heading)
	router.OPTIONS("/someOptions", options)
	router.Run()
}
