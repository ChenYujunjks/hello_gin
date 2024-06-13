package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
