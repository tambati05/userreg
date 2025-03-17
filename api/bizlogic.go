package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"userreg/dataservice" 
	"userreg/model"
	"errors"
	"fmt"
)

func RegisterUserLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error{
	return dataservice.RegisterUser(db, w, r)
}

// LoginUserLogic handles the logic for logging in a user
func LoginUserLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	// Perform basic validations here if needed

	// Call the LoginUser function from the dataservice package to handle login logic
	return dataservice.LoginUser(db, w, r)
}

// UpdateUserLogic updates a user's details by ID
func UpdateUserLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	// Extract user ID from URL parameters
	//id := r.URL.Path[len("/users/"):]
	fmt.Println("method")
	var user model.User

	// Decode the request body into the User struct
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return err
	}

	// Update the user in the database
	if err := dataservice.UpdateUser(db, user); err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return err
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
	return nil
}

// DeleteUser deletes a user by ID
func DeleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
    // Extract user ID from URL parameters
  //  id := r.URL.Path[len("/users/"):]

    // Delete user from the database
   // result, err := db.Exec("DELETE FROM user WHERE id = ?", id)
    if err != nil {
        return errors.New("failed to delete user")
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return errors.New("failed to delete user")
    }

    if rowsAffected == 0 {
        return errors.New("user not found")
    }

    // Send a success response (204 No Content)
    w.WriteHeader(http.StatusNoContent)
    return nil
}