package main

import "github.com/gin-gonic/gin"

//默认使用
func main() {
	router := gin.Default()
	goodsGroup := router.Group("/goods")
	{
		goodsGroup.GET("/list", goodsList)
		goodsGroup.GET("/login", goodsLogin)
		goodsGroup.POST("/read", goodsRead)
	}
	router.Run()
}

func goodsRead(context *gin.Context) {

}

func goodsLogin(context *gin.Context) {

}

func goodsList(context *gin.Context) {

}
