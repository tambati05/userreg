package api

import (
	"database/sql"
	"net/http"
)

// RegisterRoutes sets up the HTTP routes for user registration.
func RegisterRoutes(db *sql.DB) {
	http.HandleFunc("/update", UpdateUserHandler(db)) // Custom route name for updating user info
}
