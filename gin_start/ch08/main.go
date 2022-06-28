package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//自定义一个中间件
func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", 123456)
		//让原本该执行的逻辑继续执行
		c.Next()
		end := time.Since(t)
		fmt.Printf("耗时:%V\n", end)
		status := c.Writer.Status()
		fmt.Printf("状态", status)
	}
}
func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			if k == "X-Token" {
				token = v[0]
			} else {
				fmt.Println(k, v)
			}
		}
		if token != "bobby" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"mag": "未登录",
			})
			c.Abort()
		}
		c.Next()
	}
}
func main() {
	router := gin.Default()
	//使用logger和recovery中间件(全局使用)
	router.Use(TokenRequired())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run()
	//局部使用
	//authrized := router.Group("/goods")
	//authrized.Use(Authrized)
}

//func Authrized(context *gin.Context) {
//
//}
