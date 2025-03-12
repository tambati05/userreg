package api

import (
	"database/sql"
	"net/http"
)

func RegisterUserHandler(db *sql.DB) http.HandleFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
			return
		}
		if err := RegisterUserLogic(db, w, r); err!=nil{
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
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
