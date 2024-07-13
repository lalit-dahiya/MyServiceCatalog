package models

// User struct
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // To be replaced with encrypted hash later
}
