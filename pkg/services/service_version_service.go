package services

import "github.com/lalit-dahiya/MyServiceCatalog/api/models"

// ServiceVersionInterface to define methods on ServiceVersion struct
type ServiceVersionInterface interface {
	GetServiceVersions(serviceId string) ([]models.ServiceVersion, error)
	GetServiceVersion(id string) (models.ServiceVersion, error)
	CreateServiceVersion(service models.ServiceVersion) error
	UpdateServiceVersion(id string, service models.ServiceVersion) error
	DeleteServiceVersion(id string) error
}
