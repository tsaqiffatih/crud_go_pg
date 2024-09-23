package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/tsaqiffatih/crud_go_pg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DestroyDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"), os.Getenv("DB_TIMEZONE"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL: ", err)
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS crud_go_pg")
	if err != nil {
		log.Fatal("Failed to drop database: ", err)
	}
	fmt.Println("Database dropped successfully!")
}

func CreateDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL: ", err)
	}
	defer db.Close()

	var dbName string
	err = db.QueryRow("SELECT datname FROM pg_database WHERE datname = 'crud_go_pg'").Scan(&dbName)
	if err == sql.ErrNoRows {

		_, err = db.Exec("CREATE DATABASE crud_go_pg")
		if err != nil {
			log.Fatal("Failed to create database: ", err)
		}
		fmt.Println("Database created successfully!")
	} else if err != nil {
		log.Fatal("Error checking database: ", err)
	} else {
		fmt.Println("Database already exists!")
	}
}

func ConnectDatabase() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"), os.Getenv("DB_TIMEZONE"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("Database connected!")
}

// Function for migration database
func MigrateDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	log.Println("Database migrated successfully!")
}

// Function Seeding data to database
func SeedDatabase() {
	users := []models.User{
		{Name: "Budi Susanto", Email: "budiSusanto@mail.com"},
		{Name: "Sangkuriang Joko", Email: "sangkuriang@mail.com"},
	}

	for _, user := range users {
		result := DB.Create(&user)
		if result.Error != nil {
			log.Println("Error seeding user: ", result.Error)
		} else {
			log.Println("User seeded: ", user)
		}
	}
}
