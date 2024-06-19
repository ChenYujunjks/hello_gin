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
	// 加载HTML模板
	r.LoadHTMLGlob("views/*")

	// 将数据库实例传递给控制器
	authController := c.NewAuthController(db)

	// 设置路由
	r.POST("/login", authController.Login)
	r.POST("/register", authController.Register)
	r.GET("/users", authController.GetUsers)
	r.GET("/login", authController.ShowLoginPage)
	r.GET("/register", authController.ShowRegisterPage)

	// 启动服务器
	r.Run(":8080")
}
