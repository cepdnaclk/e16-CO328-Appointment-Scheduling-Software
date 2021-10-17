package models

import "time"

type TimeSlot struct{
	SlotId int `json:"slotId"`
	Time string `json:"time"`
	ClintName string `json:"clientName"`
	ClientEmail string `json:"clientEmail"`
	ClientRequested bool `json:"clientRequested"`
	Approved bool `json:"approved"`

}

type DayDetail struct{
	SlotList []TimeSlot `json:"slotList"`
	Date string `json:"date"`
}

type Service struct{
	ServiceID string `json:"serviceId"`
	Email string `json:"email"`
	ServiceName string `json:"serviceName"`
	ServiceDiscription string `json:"serviceDiscription"`
	ServiceType string `json:"serviceType"`
	ExpiredDay time.Time `json:"expiryDay"`

}

type ServiceDayDetail struct{
	DayDetails []DayDetail `json:"dayDetails"`
	ServiceID string `json:"serviceId"`
	Email string `json:"email"`
}



