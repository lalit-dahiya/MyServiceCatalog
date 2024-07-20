package models

// Service represents a service in my service catalog
type Service struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ServiceVersion represents a service version in my service catalog
type ServiceVersion struct {
	ID             string `json:"id"`
	Version        string `json:"version"`
	ServiceId      string `json:"serviceId"`
	NumOfDownloads int    `json:"numOfDownloads"`
}

// ServiceSummary represents a service summary in my service catalog
type ServiceSummary struct {
	Service       `json:",inline"`
	NumOfVersions int `json:"numOfVersions"`
}
