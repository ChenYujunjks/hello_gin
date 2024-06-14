package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdditionRequest struct {
	Number1 float64 `json:"number1" binding:"required"`
	Number2 float64 `json:"number2" binding:"required"`
}

func main() {
	gin.SetMode(gin.DebugMode) // 设置Gin的运行模式为DebugMode

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	r.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		// 在这里添加注册逻辑，比如保存用户信息到数据库
		c.JSON(http.StatusOK, gin.H{
			"message":  "注册成功",
			"username": username,
			"password": password,
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "admin" && password == "123456" {
			c.JSON(http.StatusOK, gin.H{
				"message": "登录成功",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "登录失败",
			})
		}
	})
	//路径参数绑定
	r.GET("user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
		})
	})
	//查询参数绑定
	r.GET("/multiply", func(c *gin.Context) {
		number1Str := c.Query("number1")
		number2Str := c.Query("number2")

		// 如果没有提供查询参数，则渲染HTML表单
		if number1Str == "" && number2Str == "" {
			c.HTML(http.StatusOK, "multiply.html", nil)
			return
		}
		// 处理提交的查询参数
		if number1Str == "" || number2Str == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "number1 and number2 are required"})
			return
		}
		number1, err1 := strconv.ParseFloat(number1Str, 64)
		number2, err2 := strconv.ParseFloat(number2Str, 64)

		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number format"})
			return
		}

		result := number1 * number2
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	r.GET("/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add.html", nil)
	})
	r.GET("/subtract", func(c *gin.Context) {
		c.HTML(http.StatusOK, "subtract.html", nil)
	})
	//form 参数绑定
	r.POST("/subtract", func(c *gin.Context) {
		number1 := c.PostForm("number1")
		number2 := c.PostForm("number2")

		if number1 == "" || number2 == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "number1 and number2 are required"})
			return
		}
		num1, err1 := strconv.ParseFloat(number1, 64)
		num2, err2 := strconv.ParseFloat(number2, 64)
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number format"})
			return
		}
		result := num1 - num2
		c.JSON(http.StatusOK, gin.H{"result": result})
	})
	//JSON 参数示例
	r.POST("/add", func(c *gin.Context) {
		number1 := c.PostForm("number1")
		number2 := c.PostForm("number2")

		if number1 == "" || number2 == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "number1 and number2 are required"})
			return
		}
		num1, err1 := strconv.ParseFloat(number1, 64)
		num2, err2 := strconv.ParseFloat(number2, 64)
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number format"})
			return
		}
		result := num1 + num2
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	r.Run(":8080")
}
