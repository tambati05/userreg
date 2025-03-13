package api

import (
	"database/sql"
	"net/http"
)

// LoginHandle handles login requests
func LoginHandle(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure that the request method is POST
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // Return a 405 Method Not Allowed response
			return
		}

		// Call LoginUserLogic to handle login logic
		if err := LoginUserLogic(db, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // Return 500 Internal Server Error on failure
		}
	}
}

// UpdateUserHandler handles PUT requests to update a user's info.
func UpdateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is PUT
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Call the logic to update user info
		if err := UpdateUserLogic(db, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
