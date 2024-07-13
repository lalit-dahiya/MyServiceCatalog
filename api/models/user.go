package models

// User struct
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // To be replaced with encrypted hash later
}

// UserInterface to define methods on User struct
type UserInterface interface {
	GetUsers() ([]User, error)
	GetUser(id string) (User, error)
	CreateUser(service User) error
	UpdateUser(id string, service User) error
	DeleteUser(id string) error
}
