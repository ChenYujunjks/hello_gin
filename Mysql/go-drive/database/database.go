package main

import (
	"database/sql"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func CreateTable(db *sql.DB) error {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS fuckers (
            id INT AUTO_INCREMENT,
            name VARCHAR(50) NOT NULL,
            age INT,
            PRIMARY KEY (id)
        );`
	_, err := db.Exec(createTableSQL)
	return err
}

func InsertUser(db *sql.DB, name string, age int) (int64, error) {
	insertSQL := `INSERT INTO fuckers (name, age) VALUES (?, ?)`
	res, err := db.Exec(insertSQL, name, age)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func GetUser(db *sql.DB, id int64) (User, error) {
	var user User
	querySQL := `SELECT id, name, age FROM fuckers WHERE id = ?`
	err := db.QueryRow(querySQL, id).Scan(&user.ID, &user.Name, &user.Age)
	return user, err
}

func UpdateUser(db *sql.DB, id int, age int) error {
	updateSQL := `UPDATE fuckers SET age = ? WHERE id = ?`
	_, err := db.Exec(updateSQL, age, id)
	return err
}

func DeleteUser(db *sql.DB, id int) error {
	deleteSQL := `DELETE FROM fuckers WHERE id = ?`
	_, err := db.Exec(deleteSQL, id)
	return err
}
