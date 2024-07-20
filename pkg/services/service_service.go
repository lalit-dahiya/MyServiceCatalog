package services

import "github.com/lalit-dahiya/MyServiceCatalog/api/models"

// ServiceInterface to define methods on Service struct
type ServiceInterface interface {
	GetServices() ([]models.ServiceSummary, error)
	SearchServices(search string) ([]models.Service, error)
	GetService(id string) (models.Service, error)
	CreateService(service models.Service) error
	UpdateService(id string, service models.Service) error
	DeleteService(id string) error
}
