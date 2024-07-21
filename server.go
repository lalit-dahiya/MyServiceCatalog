package main

import (
	"context"
	"fmt"
	"github.com/lalit-dahiya/MyServiceCatalog/api/handlers"
	"github.com/lalit-dahiya/MyServiceCatalog/api/models"
	"github.com/lalit-dahiya/MyServiceCatalog/pkg/config"
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

	cfg, err := config.LoadConfig("./config.yaml")
	if err != nil {
		panic(err)
	}

	var userService services.UserInterface
	var serviceService services.ServiceInterface
	var serviceVersionService services.ServiceVersionInterface

	ctx := context.Background()
	if cfg.Database.UseMongoDb {
		// Initialize mongoDb service
		bsonOpts := &options.BSONOptions{
			UseJSONStructTags: true,
			NilMapAsEmpty:     true,
			NilSliceAsEmpty:   true,
		}
		uri := fmt.Sprintf("mongodb://%s:%s", cfg.Database.Host, cfg.Database.Port)
		clientOpts := options.Client().ApplyURI(uri).SetBSONOptions(bsonOpts)
		mongoClient, err := mongo.Connect(ctx, clientOpts)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := mongoClient.Disconnect(ctx); err != nil {
				fmt.Println("Failed to disconnect from MongoDB:", err)
			}
		}()

		userService, err = nosql.NewUserService(ctx, mongoClient, cfg.Database.Name, cfg.Database.UserCol)
		if err != nil {
			fmt.Println("error creating user service : ", err)
			panic(err) // Handle connection or collection creation error
		}

		serviceVersionService, err = nosql.NewServiceVersionService(ctx, mongoClient, cfg.Database.Name, cfg.Database.VersionCol)
		if err != nil {
			fmt.Println("error creating version service : ", err)
			panic(err) // Handle connection or collection creation error
		}

		serviceService, err = nosql.NewServiceService(ctx, mongoClient, cfg.Database.Name, cfg.Database.ServiceCol, cfg.Database.VersionCol)
		if err != nil {
			fmt.Println("error creating service service : ", err)
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
	versionHandler := handlers.NewServiceVersionHandler(serviceVersionService)

	// Register user API handlers
	e.GET("/api/v1/users/:username", userHandler.GetUser)
	e.POST("/api/v1/users", userHandler.CreateUser)
	e.PUT("/api/v1/users/:username", userHandler.UpdateUser)
	e.DELETE("/api/v1/users/:username", userHandler.DeleteUser)

	// Register service API handlers
	e.GET("/api/v1/services", serviceHandler.GetServices)
	e.GET("/api/v1/services/search/:search", serviceHandler.SearchServices)
	e.GET("/api/v1/services/:id", serviceHandler.GetService)
	e.POST("/api/v1/services", serviceHandler.CreateService)
	e.PUT("/api/v1/services/:id", serviceHandler.UpdateService)
	e.DELETE("/api/v1/services/:id", serviceHandler.DeleteService)

	// Register service version API handlers
	e.GET("/api/v1/versions", versionHandler.GetServiceVersions)
	e.GET("/api/v1/versions/:id", versionHandler.GetServiceVersion)
	e.POST("/api/v1/versions", versionHandler.CreateServiceVersion)
	e.PUT("/api/v1/versions/:id", versionHandler.UpdateServiceVersion)
	e.DELETE("/api/v1/versions/:id", versionHandler.DeleteService)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}
