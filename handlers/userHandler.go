package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tsaqiffatih/crud_go_pg/config"
	"github.com/tsaqiffatih/crud_go_pg/models"

	"github.com/gorilla/mux"
)

// GET /users - Get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GET /users/{id} - Get a single user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// POST /users - Create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	result := config.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// PUT /users/{id} - Update an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&user)
	config.DB.Save(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DELETE /users/{id} - Delete a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user models.User
	result := config.DB.Delete(&user, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
}
