package dataservice

import(
	"database/sql"
	"encoding/json"
	"net/http"
	"userreg/model"
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

func UpdateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var user dataservice.User

	// Decode the incoming JSON request
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return errors.New("invalid request body")
	}

	// Update user details in the database
	_, err := db.Exec("UPDATE users SET password = ? WHERE username = ?", user.Password, user.Username)
	if err != nil {
		return errors.New("failed to update user info")
	}

	// Send a success response
	w.Write([]byte("User information updated successfully"))
	return nil
}

// DeleteUser deletes a user from the database by ID.
func DeleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request, id string) error {
    query := "DELETE FROM users WHERE id = ?"
    result, err := db.Exec(query, id)
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