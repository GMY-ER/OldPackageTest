package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//写一个表单验证
type LoginForm struct {
	User     string `json:"user" binding:"required,min=3,max=10"`
	Password string `json:"password" binding:"required"`
}

//登陆表单
type SignUpForm struct {
	Age        uint8  `json:"age" binding:"required"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"` //跨字段
}

func main() {
	router := gin.Default()
	router.POST("/loginJSON", loginJSON)
	router.POST("/signup", signUpForm)
	router.Run()
}

func signUpForm(c *gin.Context) {
	var signForm SignUpForm
	//如果出现错误
	if err := c.ShouldBind(&signForm); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "登陆成功",
	})
}

func loginJSON(c *gin.Context) {
	var loginForm LoginForm
	if err := c.ShouldBind(&loginForm); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}
