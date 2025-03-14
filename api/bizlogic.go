package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"userreg/dataservice"
	// Adjust the import path based on your project structure
)

// LoginUserLogic handles the logic for logging in a user
func LoginUserLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	// Perform basic validations here if needed

	// Call the LoginUser function from the dataservice package to handle login logic
	return dataservice.LoginUser(db, w, r)
}

// UpdateUserLogic updates a user's details by ID
func UpdateUserLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	// Extract user ID from URL parameters
	id := r.URL.Path[len("/users/"):]

	var user dataservice.User

	// Decode the request body into the User struct
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return err
	}

	// Update the user in the database
	if err := dataservice.UpdateUser(db, id, user); err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return err
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
	return nil
}
