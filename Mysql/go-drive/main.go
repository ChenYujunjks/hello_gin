package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数据库连接字符串
	dsn := "root:1532@tcp(127.0.0.1:3306)/mydatabase" //"username:password@tcp(127.0.0.1:3306)/mydatabase"

	// 创建数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查连接是否成功
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("连接成功！")
}
