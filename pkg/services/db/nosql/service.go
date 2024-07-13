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
)

type MongoServiceService struct {
	client *mongo.Client
	col    *mongo.Collection
	context.Context
}

func NewServiceService(ctx context.Context, client *mongo.Client, dbName, colName string) (*MongoServiceService, error) {
	col := client.Database(dbName).Collection(colName)
	return &MongoServiceService{client: client, col: col, Context: ctx}, nil
}

func (s *MongoServiceService) GetServices() ([]models.Service, error) {
	var services []models.Service
	cur, err := s.col.Find(s.Context, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get services: %w", err)
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {

		}
	}(cur, s.Context)

	for cur.Next(s.Context) {
		var service models.Service
		err := cur.Decode(&service)
		if err != nil {
			return nil, fmt.Errorf("failed to decode service: %w", err)
		}
		services = append(services, service)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate cursor: %w", err)
	}

	return services, nil
}

func (s *MongoServiceService) GetService(id string) (models.Service, error) {
	var service models.Service
	filter := bson.M{"_id": bson.M{"$eq": id}}
	err := s.col.FindOne(s.Context, filter).Decode(&service)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return models.Service{}, echo.ErrNotFound
	}
	if err != nil {
		return models.Service{}, fmt.Errorf("failed to get service: %w", err)
	}
	return service, nil
}

func (s *MongoServiceService) CreateService(newService models.Service) error {
	// Generate a unique ID
	newService.ID = primitive.NewObjectID().Hex()

	_, err := s.col.InsertOne(s.Context, newService)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}
	return nil
}

func (s *MongoServiceService) UpdateService(id string, updatedService models.Service) error {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": updatedService}

	result, err := s.col.UpdateOne(s.Context, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update service: %w", err)
	}
	if result.MatchedCount == 0 {
		return echo.ErrNotFound // Service not found
	}
	return nil
}

func (s *MongoServiceService) DeleteService(id string) error {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result, err := s.col.DeleteOne(s.Context, filter)
	if err != nil {
		return fmt.Errorf("failed to delete service: %w", err)
	}
	if result.DeletedCount == 0 {
		return echo.ErrNotFound // Service not found
	}
	return nil
}
