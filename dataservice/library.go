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
	// Assuming the user wants to change their old username to a new one, along with the password.
	
	_, err := db.Exec("UPDATE user SET new_username = ?, password = ? WHERE username = ?", user.NewUsername, user.Password, user.Username)
	if err != nil {
		return errors.New("failed to update user info") // Return an error if the update query fails.
	}

	return nil // Successfully updated the user info.
}

// DeleteUser deletes a user by username.
func DeleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request) error {

	var user model.User
	// Decode the request body into the user struct
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return err
    }

	// Delete user from the database by username.
	result, err := db.Exec("DELETE FROM user WHERE username = ?", user.Username)
	if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return err
	}

	if rowsAffected == 0 {
			http.Error(w, "User not found", http.StatusNotFound)
			return nil
	}

	// Send a success response (204 No Content)
	w.WriteHeader(http.StatusNoContent)
	return nil
}



