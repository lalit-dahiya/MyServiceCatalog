package auth

import (
	"github.com/lalit-dahiya/MyServiceCatalog/api/models"
)

type Helper struct {
	userInterface models.UserInterface
}

// Authenticate checks username and password
func (a *Helper) Authenticate(username, password string) bool {
	user, err := a.userInterface.GetUser(username)
	if err != nil {
		return false
	}
	return user.Password == password // In-memory check
}
