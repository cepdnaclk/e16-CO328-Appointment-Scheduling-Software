package models

type ServiceResponse struct{
	ServiceID string `json:"serviceId"`
	ServiceDiscription string `json:"serviceDiscription"`
	ServiceType string `json:"serviceType"`
	ServiceName string `json:"serviceName"`
}