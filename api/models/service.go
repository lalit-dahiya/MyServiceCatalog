package models

// Service represents a service in my service catalog
type Service struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ServiceInterface to define methods on Service struct
type ServiceInterface interface {
	GetServices() ([]Service, error)
	GetService(id string) (Service, error)
	CreateService(service Service) error
	UpdateService(id string, service Service) error
	DeleteService(id string) error
}
