package handlers

import (
	"fmt"
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
		fmt.Println("error getting services", err)
		return err
	}
	return c.JSON(http.StatusOK, svc)
}

// SearchServices retrieves the list of services that match the query string
func (h *ServiceHandler) SearchServices(c echo.Context) error {
	search := c.Param("search")
	svc, err := h.serviceInterface.SearchServices(search)
	if err != nil {
		fmt.Println("error getting services", err)
		return err
	}
	return c.JSON(http.StatusOK, svc)
}

// GetService retrieves a specific service by id
func (h *ServiceHandler) GetService(c echo.Context) error {
	serviceId := c.Param("id")
	service, err := h.serviceInterface.GetService(serviceId)
	if err != nil {
		fmt.Println(fmt.Sprintf("error getting service %s", serviceId), err)
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, service)
}

// CreateService creates a new service
func (h *ServiceHandler) CreateService(c echo.Context) error {
	var newService models.Service
	err := c.Bind(&newService)
	if err != nil {
		fmt.Println("error creating service", err)
		return echo.ErrBadRequest
	}
	//TODO validations
	err = h.serviceInterface.CreateService(newService)
	if err != nil {
		fmt.Println("backend error in creating service", err)
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
		fmt.Println("error updating service", err)
		return echo.ErrBadRequest
	}
	//TODO validations
	err = h.serviceInterface.UpdateService(serviceId, updatedService)
	if err != nil {
		fmt.Println(fmt.Sprintf("error updating service %s", serviceId), err)
		return echo.ErrBadRequest
	}
	return nil
}

// DeleteService deletes a service
func (h *ServiceHandler) DeleteService(c echo.Context) error {
	serviceId := c.Param("id")
	err := h.serviceInterface.DeleteService(serviceId)
	if err != nil {
		fmt.Println(fmt.Sprintf("error deleting service %s", serviceId), err)
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
