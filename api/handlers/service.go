package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/lalit-dahiya/MyServiceCatalog/api/models"
	"github.com/lalit-dahiya/MyServiceCatalog/pkg/services"
	"net/http"
)

type ServiceHandler struct {
	serviceInterface services.ServiceInterface
}

// NewServiceHandler creates a new service handler with the provided ServiceInterface
func NewServiceHandler(serviceInterface services.ServiceInterface) *ServiceHandler {
	return &ServiceHandler{
		serviceInterface: serviceInterface,
	}
}

// GetServices retrieves the list of services
func (h *ServiceHandler) GetServices(c echo.Context) error {
	svc, err := h.serviceInterface.GetServices()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, svc)
}

// GetService retrieves a specific service by id
func (h *ServiceHandler) GetService(c echo.Context) error {
	serviceId := c.Param("id")
	service, err := h.serviceInterface.GetService(serviceId)
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, service)
}

// CreateService creates a new service
func (h *ServiceHandler) CreateService(c echo.Context) error {
	var newService models.Service
	err := c.Bind(&newService)
	if err != nil {
		return echo.ErrBadRequest
	}
	//TODO validations
	err = h.serviceInterface.CreateService(newService)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusCreated, newService)
}

// UpdateService updates a service
func (h *ServiceHandler) UpdateService(c echo.Context) error {
	serviceId := c.Param("id")
	var updatedService models.Service
	err := c.Bind(&updatedService)
	if err != nil {
		return echo.ErrBadRequest
	}
	//TODO validations
	err = h.serviceInterface.UpdateService(serviceId, updatedService)
	if err != nil {
		return echo.ErrBadRequest
	}
	return echo.ErrNotFound
}

// DeleteService deletes a service
func (h *ServiceHandler) DeleteService(c echo.Context) error {
	serviceId := c.Param("id")
	err := h.serviceInterface.DeleteService(serviceId)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
