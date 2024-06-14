package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdditionRequest struct {
	Number1 float64 `json:"number1" binding:"required":`
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
		// 为简化示例，这里直接返回用户信息
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
		// 在这里添加登录验证逻辑，比如检查用户名和密码
		// 为简化示例，这里直接返回登录信息
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
	r.GET("search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		//username := c.Quer y("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
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
		var request AdditionRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := request.Number1 + request.Number2
		c.JSON(http.StatusOK, gin.H{"result", result})
	})
	fmt.Println("-----------------------------", http.StatusIMUsed) // 输出 226

	r.Run(":8080")
}
