package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/welcome", welcome)
	router.POST("/form_post", formPost)
	router.POST("/post", getPost)
	router.Run()
}

//混合的既有post又有get
func getPost(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.DefaultPostForm("message", "信息")
	c.JSON(200, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}

func formPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(200, gin.H{
		"message": message,
		"nick":    nick,
	})
}

func welcome(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "bobby")
	lastName := c.DefaultQuery("lastname", "imooc")
	c.JSON(200, gin.H{
		"first_Name": firstName,
		"last_name":  lastName,
	})
}
