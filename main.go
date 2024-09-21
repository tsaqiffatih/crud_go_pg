package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tsaqiffatih/crud_go_pg/handlers"
	"github.com/tsaqiffatih/crud_go_pg/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func destroyDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"), os.Getenv("DB_TIMEZONE"))

	// Koneksi ke PostgreSQL
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

// Fungsi untuk membuat database jika belum ada
func createDatabase() {
	// Ambil DSN dari environment variables untuk fleksibilitas
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"))

	// Koneksi ke PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL: ", err)
	}
	defer db.Close()

	// Periksa apakah database sudah ada
	var dbName string
	err = db.QueryRow("SELECT datname FROM pg_database WHERE datname = 'crud_go_pg'").Scan(&dbName)
	if err == sql.ErrNoRows {
		// Jika tidak ada, buat database
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

// Fungsi untuk menghubungkan aplikasi ke database
func connectDatabase() {
	// Ambil DSN dari environment variables untuk fleksibilitas
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

// Fungsi untuk migrasi database (membuat tabel sesuai model)
func migrateDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	log.Println("Database migrated successfully!")
}

// Fungsi untuk seeding data awal ke database
func seedDatabase() {
	users := []models.User{
		{Name: "John Doe", Email: "john@example.com"},
		{Name: "Jane Smith", Email: "jane@example.com"},
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

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	destroyDatabase()

	// Membuat database jika belum ada
	createDatabase()

	// Menghubungkan ke database
	connectDatabase()

	// Migrasi tabel database
	migrateDatabase()

	// Seeding data ke tabel users
	// seedDatabase()

	// Membagikan instance DB ke handlers
	handlers.DB = DB

	// Setup router menggunakan Gorilla Mux
	r := mux.NewRouter()

	// Routes untuk users
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// Menjalankan server di port 8000
	log.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
