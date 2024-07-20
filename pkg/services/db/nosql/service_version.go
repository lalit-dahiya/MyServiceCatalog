package nosql

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lalit-dahiya/MyServiceCatalog/api/models"
	"github.com/lalit-dahiya/MyServiceCatalog/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoServiceVersionService struct {
	client *mongo.Client
	col    *mongo.Collection
	context.Context
}

func NewServiceVersionService(ctx context.Context, client *mongo.Client, dbName, colName string) (*MongoServiceVersionService, error) {
	col := client.Database(dbName).Collection(colName)
	return &MongoServiceVersionService{client: client, col: col, Context: ctx}, nil
}

func (s *MongoServiceVersionService) GetServiceVersions(serviceId string) ([]models.ServiceVersion, error) {
	var versions []models.ServiceVersion
	cur, err := s.col.Find(s.Context, bson.M{"serviceId": bson.M{"$eq": serviceId}})
	if err != nil {
		return nil, fmt.Errorf("failed to get services: %w", err)
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {

		}
	}(cur, s.Context)

	for cur.Next(s.Context) {
		var version models.ServiceVersion
		err := cur.Decode(&version)
		if err != nil {
			return nil, fmt.Errorf("failed to decode service version: %w", err)
		}
		versions = append(versions, version)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate cursor: %w", err)
	}

	return versions, nil
}

func (s *MongoServiceVersionService) GetServiceVersion(id string) (models.ServiceVersion, error) {
	var version models.ServiceVersion
	filter := bson.M{"id": bson.M{"$eq": id}}
	err := s.col.FindOne(s.Context, filter).Decode(&version)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return models.ServiceVersion{}, echo.ErrNotFound
	}
	if err != nil {
		return models.ServiceVersion{}, fmt.Errorf("failed to get service version: %w", err)
	}
	return version, nil
}

func (s *MongoServiceVersionService) CreateServiceVersion(newService models.ServiceVersion) error {
	newService.ID = utils.GetVersionId(newService.ServiceId, newService.Version)

	_, err := s.col.InsertOne(s.Context, newService)
	if err != nil {
		return fmt.Errorf("failed to create service version: %w", err)
	}
	return nil
}

func (s *MongoServiceVersionService) UpdateServiceVersion(id string, updatedVersion models.ServiceVersion) error {
	updatedVersion.ID = id
	filter := bson.M{"id": bson.M{"$eq": id}}
	update := bson.M{"$set": updatedVersion}

	result, err := s.col.UpdateOne(s.Context, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update service version: %w", err)
	}
	if result.MatchedCount == 0 {
		return echo.ErrNotFound
	}
	return nil
}

func (s *MongoServiceVersionService) DeleteServiceVersion(id string) error {
	filter := bson.M{"id": bson.M{"$eq": id}}
	result, err := s.col.DeleteOne(s.Context, filter)
	if err != nil {
		return fmt.Errorf("failed to delete service version: %w", err)
	}
	if result.DeletedCount == 0 {
		return echo.ErrNotFound
	}
	return nil
}
