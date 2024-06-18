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

	// 创建Gin引擎
	r := gin.Default()

	// 将数据库实例传递给控制器
	authController := c.NewAuthController(db)

	// 设置路由
	r.POST("/login", authController.Login)

	// 启动服务器
	r.Run(":8080")
}
