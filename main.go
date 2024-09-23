package main

import (
	"log"
	"net/http"

	"github.com/tsaqiffatih/crud_go_pg/config"
	"github.com/tsaqiffatih/crud_go_pg/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.DestroyDatabase()

	config.CreateDatabase()

	// Connecting to database
	config.ConnectDatabase()

	// Database table migration
	config.MigrateDatabase()

	// Seeding data into tabel users
	config.SeedDatabase()

	// Setup router using Gorilla Mux
	r := mux.NewRouter()

	// Routes CRUD for users
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// Menjalankan server di port 8000
	log.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
