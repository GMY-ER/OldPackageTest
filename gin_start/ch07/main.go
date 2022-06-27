package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
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

func InitTrans(locale string) (err error) {
	//修改gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个语言参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		trans, ok := uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
	}
	return
}

var trans ut.Translator

func main() {
	//想计算方法的运行时长，如果记时，代码侵入性很强
	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化错误")
		return
	}
	router := gin.Default()
	router.POST("/loginJSON", loginJSON)
	router.POST("/signup", signUpForm)
	router.Run()
}

func signUpForm(c *gin.Context) {
	var signForm SignUpForm
	//如果出现错误
	if err := c.ShouldBind(&signForm); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs.Translate(trans),
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
