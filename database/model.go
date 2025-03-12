package database

import(
	"database/sql"
	"encoding/json"
	"net/http"
	model "userreg/module"
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