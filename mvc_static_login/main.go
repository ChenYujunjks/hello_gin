package main

import (
	"log"
	c "login_mvc/controllers"
	"login_mvc/models"

	"github.com/gin-gonic/gin"

	//选择轻量级数据库而不是"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// SQLite连接配置
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 迁移数据库
	db.AutoMigrate(&models.User{})

	r := gin.Default()
	//首先加载templates目录下面的所有模版文件，模版文件扩展名随意
	r.LoadHTMLGlob("views/pages/*")
	//  /assets/images/1.jpg 这个url文件，存储在/public/images/1.jpg
	r.Static("/assets", "views/static")
	// 将数据库实例传递给控制器
	authController := c.NewAuthController(db)

	// 设置路由
	r.POST("/login", authController.Login)
	r.POST("/register", authController.Register)
	r.GET("/users", authController.GetUsers)
	r.GET("/login", authController.ShowLoginPage)
	r.GET("/register", authController.ShowRegisterPage)
	r.GET("/", authController.ShowIndexPage)
	// 启动服务器
	r.Run(":8080")
}
