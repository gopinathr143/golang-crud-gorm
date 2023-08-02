package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// CreateHandler handles the creation of a new user record.
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO users (name) VALUES (?)", user.Name)
	if err != nil {
		http.Error(w, "Error inserting user data", http.StatusInternalServerError)
		return
	}

	user.ID, err = result.LastInsertId()
	if err != nil {
		http.Error(w, "Error getting last insert ID", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// ReadHandler handles reading user records.
func ReadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(w, "Error querying user data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			http.Error(w, "Error scanning user data", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UpdateHandler handles updating user records.
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE users SET name = ? WHERE id = ?", user.Name, user.ID)
	if err != nil {
		http.Error(w, "Error updating user data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteHandler handles deleting user records.
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM users WHERE id = ?", user.ID)
	if err != nil {
		http.Error(w, "Error deleting user data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
