package model

type User struct {
	Username    string `json:"username"`
	NewUsername string `json:"new_username"` 
	Password    string `json:"password"`
}
