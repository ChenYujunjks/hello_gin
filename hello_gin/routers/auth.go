package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
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

}
