package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

// RegisterRoutes sets up the HTTP routes for user registration.

func RegisterRoutes(db *sql.DB) {
	http.HandleFunc("/register", RegisterUserHandler(db))
  	// Registering the route for user login (POST)
	http.HandleFunc("/login", LoginHandle(db))
	http.HandleFunc("/update", UpdateUserHandler(db)) // Custom route name for updating user info
	http.HandleFunc("/delete", DeleteUserHandler(db)) // Route for deleting a user
	fmt.Println("Routes are successfully registered.")

}
