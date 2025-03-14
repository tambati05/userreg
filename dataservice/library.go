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

func UpdateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
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