package main

import (
	"context"
	"fmt"
	"github.com/lalit-dahiya/MyServiceCatalog/api/handlers"
	"github.com/lalit-dahiya/MyServiceCatalog/api/models"
	"github.com/lalit-dahiya/MyServiceCatalog/pkg/services"
	"github.com/lalit-dahiya/MyServiceCatalog/pkg/services/db/nosql"
	"github.com/lalit-dahiya/MyServiceCatalog/pkg/services/inmemory"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var userService services.UserInterface
	var serviceService services.ServiceInterface

	ctx := context.Background()
	useMongoDb := false // Ideally fetch from runtime config
	if useMongoDb {
		// Initialize mongoDb service
		bsonOpts := &options.BSONOptions{
			UseJSONStructTags: true,
			NilMapAsEmpty:     true,
			NilSliceAsEmpty:   true,
		}
		clientOpts := options.Client().ApplyURI("mongodb:// localhost:27000").SetBSONOptions(bsonOpts)
		mongoClient, err := mongo.Connect(ctx, clientOpts)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := mongoClient.Disconnect(ctx); err != nil {
				fmt.Println("Failed to disconnect from MongoDB:", err)
			}
		}()

		userService, err = nosql.NewUserService(ctx, mongoClient, "MyServiceCatalog", "users")
		if err != nil {
			panic(err) // Handle connection or collection creation error
		}

		serviceService, err = nosql.NewServiceService(ctx, mongoClient, "MyServiceCatalog", "services")
		if err != nil {
			panic(err) // Handle connection or collection creation error
		}
	} else {
		// Initialize in-memory service
		userService = &inmemory.User{
			Users: map[string]models.User{
				"user1": {"user1", "user1", "password1"},
				"user2": {"user2", "user2", "password2"},
			},
		}
		serviceService = &inmemory.Service{
			Services: []models.Service{
				{"1", "Service 1", "Description for service 1"},
				{"2", "Service 2", "Description for service 2"},
			},
		}
	}

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
	e.GET("/services/:id", serviceHandler.GetService)
	e.POST("/services", serviceHandler.CreateService)
	e.PUT("/services/:id", serviceHandler.UpdateService)
	e.DELETE("/services/:id", serviceHandler.DeleteService)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
