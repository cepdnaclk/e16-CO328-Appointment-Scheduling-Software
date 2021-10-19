package models

type Time struct{
	Hh string `json:"hh"`
	Mm string `json:"mm"`
	Ss string `json:"ss"`
	A string `json:"a"`
}

type RequestTimeSlot struct{
	StartTime Time `json:"startTime"`
	EndTime Time `json:"endTime"`
	Id int `json:"id"`
}

type AddNewServiceRequest struct{
	TimeSlots []RequestTimeSlot `json:"timeSlots"`
	ServiceName string `json:"serviceName"`
	ServiceDiscription string `json:"serviceDiscription"`
	ServiceType string `json:"serviceType"`
}

type ServiceIdRequsest struct{
	ServiceID string `json:"serviceId"`
}

type UpdateRequest struct{
	ServiceID string `json:"serviceId"`
	SlotId int `json:"slotId"`
	Date string `json:"date"`
}