package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"usereg/database" // Adjust the import path based on your project structure
)

// UpdateUserLogic updates a user's details by ID
func UpdateUserLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	// Extract user ID from URL parameters
	id := r.URL.Path[len("/users/"):]

	var user database.User

	// Decode the request body into the User struct
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return err
	}

	// Update the user in the database
	if err := database.UpdateUser(db, id, user); err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return err
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
	return nil
}
