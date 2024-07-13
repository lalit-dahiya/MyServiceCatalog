package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/lalit-dahiya/MyServiceCatalog/api/models"
	"net/http"
)

type UserHandler struct {
	userInterface models.UserInterface
}

// GetUsers retrieves the list of users
func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.userInterface.GetUsers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

// GetUser retrieves a specific user by id
func (h *UserHandler) GetUser(c echo.Context) error {
	userId := c.Param("id")
	user, err := h.userInterface.GetUser(userId)
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c echo.Context) error {
	var newUser models.User
	err := c.Bind(&newUser)
	if err != nil {
		return echo.ErrBadRequest
	}
	// validations
	err = h.userInterface.CreateUser(newUser)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusCreated, newUser)
}

// UpdateUser updates a user
func (h *UserHandler) UpdateUser(c echo.Context) error {
	userId := c.Param("id")
	var updatedUser models.User
	err := c.Bind(&updatedUser)
	if err != nil {
		return echo.ErrBadRequest
	}
	// validations
	err = h.userInterface.UpdateUser(userId, updatedUser)
	if err != nil {
		return echo.ErrBadRequest
	}
	return echo.ErrNotFound
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(c echo.Context) error {
	userId := c.Param("id")
	err := h.userInterface.DeleteUser(userId)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
