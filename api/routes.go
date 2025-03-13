package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

// RegisterRoutes sets up the HTTP routes for user registration.

func RegisterRoutes(db *sql.DB) {
	// Registering the route for user login (POST)
	http.HandleFunc("/login", LoginHandle(db))

	http.HandleFunc("/update", UpdateUserHandler(db))

	fmt.Println("Routes are successfully registered.")
}
