package services

import "github.com/lalit-dahiya/MyServiceCatalog/api/models"

// UserInterface to define methods on User struct
type UserInterface interface {
	GetUser(id string) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(id string, user models.User) error
	DeleteUser(id string) error
}
