package models

type ServiceResponse struct{
	ServiceID string `json:"serviceId"`
	ServiceDiscription string `json:"serviceDiscription"`
	ServiceType string `json:"serviceType"`
	ServiceName string `json:"serviceName"`
}

type ClientServiceResponse struct{
	ServiceOwnerEmail string `json:"serviceOwnerEmail"`
	ServiceID string `json:"serviceId"`
	ServiceDiscription string `json:"serviceDiscription"`
	ServiceType string `json:"serviceType"`
	ServiceName string `json:"serviceName"`
	OwnerFristName string  `json:"ownerFristName"`
	OwnerLastName string  `json:"ownerLastName"`
}