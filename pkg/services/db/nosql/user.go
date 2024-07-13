package nosql

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lalit-dahiya/MyServiceCatalog/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MongoUserService struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewUserService(ctx context.Context, client *mongo.Client, dbName, colName string) (*MongoUserService, error) {
	col := client.Database(dbName).Collection(colName)
	return &MongoUserService{client: client, col: col}, nil
}

func (s *MongoUserService) GetUser(username string) (models.User, error) {
	var user models.User
	filter := bson.M{"username": username}
	err := s.col.FindOne(context.Background(), filter).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return models.User{}, echo.ErrNotFound
	}
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (s *MongoUserService) CreateUser(newUser models.User) error {
	// Generate a unique ID
	newUser.ID = primitive.NewObjectID().Hex()

	log.Printf("Creating user %s with id %s", newUser.Username, newUser.ID)
	_, err := s.col.InsertOne(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	log.Printf("User %s created successfully with id %s", newUser.Username, newUser.ID)
	return nil
}

func (s *MongoUserService) UpdateUser(username string, updatedUser models.User) error {
	filter := bson.M{"username": username}
	update := bson.M{"$set": updatedUser}

	result, err := s.col.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	if result.MatchedCount == 0 {
		return echo.ErrNotFound // User not found
	}
	return nil
}

func (s *MongoUserService) DeleteUser(username string) error {
	filter := bson.M{"username": username}
	result, err := s.col.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	if result.DeletedCount == 0 {
		return echo.ErrNotFound // User not found
	}
	return nil
}
