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

type ClientRequestedServicesResponse struct{
	ServiceOwnerEmail string `json:"serviceOwnerEmail"`
	ServiceID string `json:"serviceId"`
	SlotId int `json:"slotId"`
	Approved bool `json:"approved"`
	Time string `json:"time"`
	ServiceName string `json:"serviceName"`
	ServiceDiscription string `json:"serviceDiscription"`
	Date string `json:"date"`
}

type ClientTimeSlot struct{
	SlotId int `json:"slotId"`
	Time string `json:"time"`
	ClientRequested bool `json:"clientRequested"`
}

type ClientDayDetailResponse struct{
	SlotList []ClientTimeSlot `json:"slotList"`
	Date string `json:"date"`
}