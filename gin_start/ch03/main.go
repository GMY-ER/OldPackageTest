package main

import "github.com/gin-gonic/gin"

//默认使用
func main() {
	router := gin.Default()
	goodsGroup := router.Group("/goods")
	{
		//goodsGroup.GET("", goodsList)
		//goodsGroup.GET("", goodsLogin)
		//goodsGroup.GET("", goodsRead)
		//goodsGroup.GET("/1", goodsDetail) 设置一个变量
		goodsGroup.GET("/:id/:action", goodsDetail)
	}
	router.Run()
}

func goodsDetail(context *gin.Context) {
	id := context.Param("id")
	action := context.Param("action")
	context.JSON(200, gin.H{
		"id":     id,
		"aciton": action,
	})
}

func goodsRead(context *gin.Context) {
	context.JSON(200, gin.H{
		"name": "goodsRead",
	})
}

func goodsLogin(context *gin.Context) {

}

func goodsList(context *gin.Context) {

}
