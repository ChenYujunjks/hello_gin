package main

import (
	"log"

	c "hello_gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode) // 设置Gin的运行模式为DebugMode

	r := gin.Default()
	r.Use(StatCost())
	r.LoadHTMLGlob("views/*")

	// 注册路由
	r.GET("/multiply", c.GetMultiply)
	r.GET("/add", c.GetAdd)
	r.GET("/subtract", c.GetSubtract)
	r.POST("/subtract", c.PostSubtract)
	r.POST("/add", c.PostAdd)
	//auth
	r.GET("/upload", c.GetUpload)
	r.POST("/upload", c.PostUpload)
	r.GET("/register", c.GetRegister)
	r.POST("/register", c.PostRegister)
	r.GET("/login", c.GetLogin)
	r.POST("/login", c.PostLogin)
	r.GET("/user/:id", c.GetUser)

	r.GET("/old-page", c.RedirectToAdd)
	r.GET("/slice-any", c.GetSliceAny)
	r.GET("/slice-struct", c.GetSliceStruct)
	r.GET("/number", c.GetNumber)
	r.GET("/string", c.GetString)
	r.GET("/map", c.GetMap)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Server Run Failed:", err)
	}
}
