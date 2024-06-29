package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:1532@tcp(127.0.0.1:3306)/mydatabase"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	fmt.Println("Successfully connected to the database")

	// Create table
	if err := CreateTable(db); err != nil {
		log.Fatal(err)
	}
	// Insert user
	lastInsertID, err := InsertUser(db, "Bruce", 21)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted new user with ID %d\n", lastInsertID)
	// Get user
	user, err := GetUser(db, lastInsertID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User details: ID=%d, Name=%s, Age=%d\n", user.ID, user.Name, user.Age)
	/*
	   // Update user

	   	if err := UpdateUser(db, user.ID, 26); err != nil {
	   		log.Fatal(err)
	   	}
	   	fmt.Printf("Updated user with ID %d\n", user.ID)
	   	// Delete user

	   		if err := DeleteUser(db, user.ID); err != nil {
	   			log.Fatal(err)
	   		}
	   		fmt.Printf("Deleted user with ID %d\n", user.ID)
	*/
}
