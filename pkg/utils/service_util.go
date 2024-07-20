package utils

import "github.com/lalit-dahiya/MyServiceCatalog/api/models"

func ConvertToSummary(service models.Service, versions int) models.ServiceSummary {
	return models.ServiceSummary{
		Service:       service,
		NumOfVersions: versions,
	}
}

func GetVersionId(serviceId, version string) string {
	return serviceId + "_" + version
}
