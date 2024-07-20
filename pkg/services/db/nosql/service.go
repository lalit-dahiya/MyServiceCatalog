package nosql

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lalit-dahiya/MyServiceCatalog/api/models"
	"github.com/lalit-dahiya/MyServiceCatalog/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"regexp"
)

type MongoServiceService struct {
	client     *mongo.Client
	serviceCol *mongo.Collection
	versionCol *mongo.Collection
	context.Context
}

func NewServiceService(ctx context.Context, client *mongo.Client, dbName, serviceColName, versionColName string) (*MongoServiceService, error) {
	serviceCol := client.Database(dbName).Collection(serviceColName)
	versionCol := client.Database(dbName).Collection(versionColName)
	return &MongoServiceService{client: client, serviceCol: serviceCol, versionCol: versionCol, Context: ctx}, nil
}

func (s *MongoServiceService) GetServices() ([]models.ServiceSummary, error) {
	var summaries []models.ServiceSummary
	cur, err := s.serviceCol.Find(s.Context, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get services: %w", err)
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			fmt.Println("failed to close cursor", err)
		}
	}(cur, s.Context)

	for cur.Next(s.Context) {
		var service models.Service
		err := cur.Decode(&service)
		if err != nil {
			return nil, fmt.Errorf("failed to decode service: %w", err)
		}

		//get number of versions
		filter := bson.M{"serviceId": service.ID}
		count, err := s.versionCol.CountDocuments(context.Background(), filter)
		if err != nil {
			fmt.Println(fmt.Sprintf("failed to find service versions for id %s", service.ID), err)
		}
		summary := utils.ConvertToSummary(service, int(count))
		summaries = append(summaries, summary)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate cursor: %w", err)
	}

	return summaries, nil
}

func (s *MongoServiceService) GetService(id string) (models.Service, error) {
	var service models.Service
	filter := bson.M{"id": bson.M{"$eq": id}}
	err := s.serviceCol.FindOne(s.Context, filter).Decode(&service)
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

	_, err := s.serviceCol.InsertOne(s.Context, newService)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}
	return nil
}

func (s *MongoServiceService) UpdateService(id string, updatedService models.Service) error {
	updatedService.ID = id
	filter := bson.M{"id": bson.M{"$eq": id}}
	update := bson.M{"$set": updatedService}

	result, err := s.serviceCol.UpdateOne(s.Context, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update service: %w", err)
	}
	if result.MatchedCount == 0 {
		fmt.Println("No service found with ID:", id)
		return echo.ErrNotFound
	} else if result.ModifiedCount == 1 {
		fmt.Println("Service updated successfully!")
	} else {
		fmt.Println("Unexpected update result:", result)
	}
	return nil
}

func (s *MongoServiceService) DeleteService(id string) error {
	filter := bson.M{"id": bson.M{"$eq": id}}
	result, err := s.serviceCol.DeleteOne(s.Context, filter)
	if err != nil {
		return fmt.Errorf("failed to delete service: %w", err)
	}
	if result.DeletedCount == 0 {
		return echo.ErrNotFound // Service not found
	}
	return nil
}

func (s *MongoServiceService) SearchServices(search string) ([]models.Service, error) {
	regex, err := regexp.Compile(".*" + search + ".*")
	if err != nil {
		return nil, fmt.Errorf("failed to compile regex: %w", err)
	}
	criteria := bson.M{"name": bson.M{"$regex": regex.String()}}
	var services []models.Service
	cur, err := s.serviceCol.Find(s.Context, criteria)
	if err != nil {
		return nil, fmt.Errorf("failed to search services: %w", err)
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			fmt.Println("failed to close cursor", err)
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
