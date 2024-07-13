package main

import (
	"github.com/lalit-dahiya/MyServiceCatalog/api/handlers"
	"github.com/lalit-dahiya/MyServiceCatalog/api/models"
	"github.com/lalit-dahiya/MyServiceCatalog/pkg/services/inmemory"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Initialize in-memory service
	userService := &inmemory.User{
		Users: map[string]models.User{
			"user1": {"user1", "user1", "password1"},
			"user2": {"user2", "user2", "password2"},
		},
	}
	serviceService := &inmemory.Service{
		Services: []models.Service{
			{"1", "Service 1", "Description for service 1"},
			{"2", "Service 2", "Description for service 2"},
		},
	}

	//TODO: Database connectivity

	// Create handlers with the initialized services
	userHandler := handlers.NewUserHandler(userService)
	serviceHandler := handlers.NewServiceHandler(serviceService)

	// Register user API handlers
	e.GET("/users/:username", userHandler.GetUser)
	e.POST("/users", userHandler.CreateUser)
	e.PUT("/users/:username", userHandler.UpdateUser)
	e.DELETE("/users/:username", userHandler.DeleteUser)

	// Register service API handlers
	e.GET("/services", serviceHandler.GetServices)
	e.GET("/services/:username", serviceHandler.GetService)
	e.POST("/services", serviceHandler.CreateService)
	e.PUT("/services/:username", serviceHandler.UpdateService)
	e.DELETE("/services/:username", serviceHandler.DeleteService)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
