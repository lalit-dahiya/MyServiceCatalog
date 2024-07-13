package inmemory

import (
	"github.com/labstack/echo/v4"
	"github.com/lalit-dahiya/MyServiceCatalog/api/models"
)

type Service struct {
	Services []models.Service
}

func (s *Service) GetServices() ([]models.Service, error) {
	return s.Services, nil
}

func (s *Service) GetService(id string) (models.Service, error) {
	for _, service := range s.Services {
		if service.ID == id {
			return service, nil
		}
	}
	return models.Service{}, echo.ErrNotFound
}

func (s *Service) CreateService(service models.Service) error {

	s.Services = append(s.Services, service)
	return nil
}

func (s *Service) UpdateService(id string, service models.Service) error {
	for i, svc := range s.Services {
		if svc.ID == id {
			s.Services[i] = service
			return nil
		}
	}
	return echo.ErrNotFound
}

func (s *Service) DeleteService(id string) error {
	newServices := []models.Service{}
	for _, service := range s.Services {
		if service.ID != id {
			newServices = append(newServices, service)
		}
	}
	if len(newServices) == len(s.Services) {
		return echo.ErrNotFound
	}
	s.Services = newServices
	return nil
}
