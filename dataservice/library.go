package dataservice

import (
	"database/sql"  // Importing the SQL package to interact with the database
	"encoding/json" // Importing the JSON package to parse request body and send responses
	"errors"        // Importing errors package to handle error responses
	"net/http"      // Importing the HTTP package to handle web requests
	"userreg/model" // Importing the model package to use the User struct
)

// LoginUser handles login logic by verifying the username and password
func LoginUser(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var user model.User // Declare a variable to hold the login credentials

	// Decode the JSON request body into the user struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return errors.New("invalid request body") // Return error if JSON decoding fails
	}

	// Retrieve the stored password for the user from the database
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", user.Username).Scan(&storedPassword)
	if err != nil {
		return errors.New("user not found") // Return an error if the user doesn't exist
	}

	// Compare the stored password with the provided password
	if storedPassword != user.Password {
		return errors.New("invalid credentials") // Return an error if the passwords don't match
	}

	// Send a success message if login is successful
	w.Write([]byte("Login successful"))
	return nil
}
