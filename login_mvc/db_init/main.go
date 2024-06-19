package main

import (
	"log"
	"login_mvc/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 连接到SQLite数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 迁移数据库
	db.AutoMigrate(&models.User{})

	// 创建三个用户
	users := []models.User{
		{Username: "user1", Password: "password1"},
		{Username: "user2", Password: "password2"},
		{Username: "user3", Password: "password3"},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Failed to create user %s: %v", user.Username, err)
		}
	}

	log.Println("Database initialized with 3 users")
}
