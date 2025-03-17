package dataservice


import (
	"database/sql"  // Importing the SQL package to interact with the database
	"encoding/json" // Importing the JSON package to parse request body and send responses
	"errors"        // Importing errors package to handle error responses
	"net/http"      // Importing the HTTP package to handle web requests
	"userreg/model" // Importing the model package to use the User struct
)

func RegisterUser(db *sql.DB, w http.ResponseWriter, r *http.Request) error{
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err!=nil{
		return err
	}

	query := "INSERT INTO user(username, password) VALUES(?,?)"
	_, err := db.Exec(query, user.Username, user.Password)
	if err!=nil{
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(user)
}

// LoginUser handles login logic by verifying the username and password
func LoginUser(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var user model.User // Declare a variable to hold the login credentials

	// Decode the JSON request body into the user struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return errors.New("invalid request body") // Return error if JSON decoding fails
	}

	// Retrieve the stored password for the user from the database
	var storedPassword string
	err := db.QueryRow("SELECT password FROM user WHERE username = ?", user.Username).Scan(&storedPassword)
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

func UpdateUser(db *sql.DB, user model.User) error {
	//var user model.User

	// Decode the incoming JSON request
	/*
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return errors.New("invalid request body")
	}*/

	// Update user details in the database
	_, err := db.Exec("UPDATE user SET password = ? WHERE username = ?", user.Password, user.Username)
	if err != nil {
		return errors.New("failed to update user info")
	}

	// Send a success response
	//w.Write([]byte("User information updated successfully"))
	return nil
}

// DeleteUser deletes a user from the database by ID.
func DeleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request, id string) error {
    query := "DELETE FROM user WHERE id = ?"
    result, err := db.Exec(query, username)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        http.Error(w, "User not found", http.StatusNotFound)
        return nil // User not found, but we don't return an error to the caller.
    }

    w.WriteHeader(http.StatusNoContent)
    return nil
}




