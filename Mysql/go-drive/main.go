package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// 数据库连接字符串：用户名:密码@tcp(主机:端口)/数据库名
	dsn := "root:1532@tcp(127.0.0.1:3306)/mydatabase" //"username:password@tcp(127.0.0.1:3306)/testdb"

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 测试数据库连接
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database")

	// 插入数据
	insertSQL := `INSERT INTO fuckers (name, age) VALUES (?, ?)`
	res, err := db.Exec(insertSQL, "negro", 29)
	if err != nil {
		log.Fatal(err)
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted new fuckers with ID %d\n", lastInsertID)

	var user1 User
	querySQL := `SELECT id, name, age FROM fuckers WHERE id = ?`
	if err := db.QueryRow(querySQL, lastInsertID).Scan(&user1.ID, &user1.Name, &user1.Age); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User details: ID=%d, Name=%s, Age=%d\n", user1.ID, user1.Name, user1.Age)
	// 查询所有用户
	querySQL1 := `SELECT id, name, age FROM fuckers`
	rows1, err := db.Query(querySQL1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows1.Close()

	// 创建一个切片来存储所有用户
	var users []User

	// 遍历结果集
	for rows1.Next() {
		var user User
		if err := rows1.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	// 检查 for 循环外的错误
	if err := rows1.Err(); err != nil {
		log.Fatal(err)
	}

	// 输出所有用户信息
	fmt.Println(users)

	fmt.Println("-------------------------------------------------------------------")

	// 执行 DESCRIBE users 命令
	describeSQL2 := `DESCRIBE users`
	rows2, err := db.Query(describeSQL2)
	if err != nil {
		log.Fatal(err)
	}
	defer rows2.Close()

	// 处理结果
	var field, dataType, null, key string
	var defaultValue, extra sql.NullString

	fmt.Println("Table structure of 'users':")
	for rows2.Next() {
		err := rows2.Scan(&field, &dataType, &null, &key, &defaultValue, &extra)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Field: %s, Type: %s, Null: %s, Key: %s, Default: %v, Extra: %v\n", field, dataType, null, key, defaultValue, extra)
	}

	// 检查是否有错误
	if err := rows2.Err(); err != nil {
		log.Fatal(err)
	}
}
