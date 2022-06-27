package main

import "github.com/gin-gonic/gin"

//proto 无法做东西没有配全
func main() {
	router := gin.Default()
	router.GET("/moreJSON", moreJSON)
	router.Run()
}

func moreJSON(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "gmy"
	msg.Message = "hello shopee"
	msg.Number = 23
	c.JSON(200, msg)
}
