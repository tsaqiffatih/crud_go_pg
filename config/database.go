package config

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateDatabase() {
	// Koneksi ke PostgreSQL tanpa menyertakan nama database
	dsn := "host=localhost user=postgres password=yourpassword port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL: ", err)
	}
	defer db.Close()

	// Perintah untuk membuat database
	_, err = db.Exec("CREATE DATABASE yourdb")
	if err != nil {
		log.Fatal("Failed to create database: ", err)
	}
	fmt.Println("Database created successfully!")
}
