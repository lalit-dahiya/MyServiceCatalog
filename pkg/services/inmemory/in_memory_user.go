package inmemory

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lalit-dahiya/MyServiceCatalog/api/models"
)

type User struct {
	Users map[string]models.User
}

func (s *User) GetUser(username string) (models.User, error) {
	user, ok := s.Users[username]
	if !ok {
		return models.User{}, echo.ErrNotFound
	}
	return user, nil
}

func (s *User) CreateUser(newUser models.User) error {
	// Check for existing username
	if _, exists := s.Users[newUser.Username]; exists {
		return echo.ErrConflict
	}
	newUser.ID = fmt.Sprintf("user-%d", len(s.Users)+1)
	s.Users[newUser.Username] = newUser
	return nil
}

func (s *User) UpdateUser(username string, updatedUser models.User) error {
	// Check if the user exists
	currentUser, ok := s.Users[username]
	if !ok {
		return echo.ErrNotFound
	}

	// Update user data
	currentUser.Username = updatedUser.Username
	currentUser.Password = updatedUser.Password

	s.Users[username] = currentUser
	return nil
}

func (s *User) DeleteUser(username string) error {
	_, ok := s.Users[username]
	if !ok {
		return echo.ErrNotFound
	}
	delete(s.Users, username)
	return nil
}
