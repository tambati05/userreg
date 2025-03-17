package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"userreg/dataservice" 
	"userreg/model"
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

func UpdateUserLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var user model.User

	// Decode the request body into the User struct
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Reject unknown fields in the input
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest) // Improved error message
		return err
	}

	// Update the user in the database
	if err := dataservice.UpdateUser(db, user); err != nil {
		http.Error(w, "Error updating user: "+err.Error(), http.StatusInternalServerError) // Include detailed error info
		return err
	}

	// Respond with the updated user info (or a success message if not needed)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully")) 	
	return nil
}

func DeleteUserLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.DeleteUser(db, w, r) 
   }