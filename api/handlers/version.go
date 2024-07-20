package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lalit-dahiya/MyServiceCatalog/api/models"
	"github.com/lalit-dahiya/MyServiceCatalog/pkg/services"
	"net/http"
)

type ServiceVersionHandler struct {
	versionInterface services.ServiceVersionInterface
}

// NewServiceVersionHandler creates a new service version handler with the provided ServiceVersionInterface
func NewServiceVersionHandler(versionInterface services.ServiceVersionInterface) *ServiceVersionHandler {
	return &ServiceVersionHandler{
		versionInterface: versionInterface,
	}
}

// GetServiceVersions retrieves the list of service versions given a serviceId
func (h *ServiceVersionHandler) GetServiceVersions(c echo.Context) error {
	serviceId := c.QueryParam("serviceId")
	svc, err := h.versionInterface.GetServiceVersions(serviceId)
	if err != nil {
		fmt.Println("error getting service versions", err)
		return err
	}
	return c.JSON(http.StatusOK, svc)
}

// GetServiceVersion retrieves a specific version by id
func (h *ServiceVersionHandler) GetServiceVersion(c echo.Context) error {
	versionId := c.Param("id")
	service, err := h.versionInterface.GetServiceVersion(versionId)
	if err != nil {
		fmt.Println(fmt.Sprintf("error getting service version %s", versionId), err)
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, service)
}

// CreateServiceVersion creates a new service version
func (h *ServiceVersionHandler) CreateServiceVersion(c echo.Context) error {
	var newServiceVersion models.ServiceVersion
	err := c.Bind(&newServiceVersion)
	if err != nil {
		fmt.Println("error creating service version", err)
		return echo.ErrBadRequest
	}
	//TODO validations
	err = h.versionInterface.CreateServiceVersion(newServiceVersion)
	if err != nil {
		fmt.Println("backend error in creating service version", err)
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusCreated, newServiceVersion)
}

// UpdateServiceVersion updates a service version
func (h *ServiceVersionHandler) UpdateServiceVersion(c echo.Context) error {
	versionId := c.Param("id")
	var updatedVersion models.ServiceVersion
	err := c.Bind(&updatedVersion)
	if err != nil {
		fmt.Println("error updating service version", err)
		return echo.ErrBadRequest
	}
	//TODO validations
	err = h.versionInterface.UpdateServiceVersion(versionId, updatedVersion)
	if err != nil {
		fmt.Println(fmt.Sprintf("error updating service version %s", versionId), err)
		return echo.ErrBadRequest
	}
	return nil
}

// DeleteService deletes a service
func (h *ServiceVersionHandler) DeleteService(c echo.Context) error {
	versionId := c.Param("id")
	err := h.versionInterface.DeleteServiceVersion(versionId)
	if err != nil {
		fmt.Println(fmt.Sprintf("error deleting service version %s", versionId), err)
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
