package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Rice struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:1532@tcp(127.0.0.1:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Database connection successful:", db)
	db.AutoMigrate(&Rice{})
	//fmt.Println(&Product{})
	prodd := Rice{Code: "sss", Price: 455}
	prod2 := Rice{Code: "third one", Price: 688}
	db.Create(&prodd)
	db.Create(&prod2)
}
