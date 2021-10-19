package controllers

import (
	"appoiment-backend/database"
	"appoiment-backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AddNewService(c *fiber.Ctx)  error{
	data:=new(models.AddNewServiceRequest)
	service:=new(models.Service)
	serviceDayDatails:=new(models.ServiceDayDetail)
	var dayDetail []models.DayDetail
	var timeSlotList []models.TimeSlot
	var noOfDays int

	if err := c.BodyParser(data); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":"Required arguments not found"})
	}

	service.Email=string(c.Request().Header.Peek("user_email"))
	service.ServiceType=data.ServiceType
	service.ServiceDiscription=data.ServiceDiscription
	service.ServiceName=data.ServiceName
	service.ServiceID=time.Now().String()+service.Email+data.ServiceName

	serviceDayDatails.ServiceID=service.ServiceID
	serviceDayDatails.Email=service.Email

	today := time.Now()
	if data.ServiceType=="Once" {
		service.ExpiredDay = today.Add(24 * time.Hour*2)
		noOfDays=1
	}else if data.ServiceType=="Monthly"{
		service.ExpiredDay = today.Add(24 * time.Hour*31)
		noOfDays=30
	}else{
		service.ExpiredDay = today.Add(24 * time.Hour*8)
		noOfDays=7
	}

	for i := 0; i < len(data.TimeSlots); i++ {
		timeSlotList=append(timeSlotList,
			models.TimeSlot{Time: data.TimeSlots[i].StartTime.Hh+":"+data.TimeSlots[i].StartTime.Mm+":"+data.TimeSlots[i].StartTime.Ss+":"+data.TimeSlots[i].StartTime.A+"-"+data.TimeSlots[i].EndTime.Hh+":"+data.TimeSlots[i].EndTime.Mm+":"+data.TimeSlots[i].EndTime.Ss+":"+data.TimeSlots[i].EndTime.A,SlotId: data.TimeSlots[i].Id,ClintName: "",ClientEmail: "",ClientRequested: false,Approved: false,
		})
	}

	for i := 0; i < noOfDays; i++ {
		dayDetail=append(dayDetail, models.DayDetail{
			Date: today.Add(24 * time.Hour *time.Duration(i+1)).String(),
			SlotList: timeSlotList,
		})
	}
	if database.AddNewService(service,serviceDayDatails) {
		c.Status(200)
		return c.JSON(fiber.Map{"message":"Done"})
	}else {
		c.Status(409)
		return c.JSON(fiber.Map{"message":"Error Occured"})
	}

}


func GetAllServices(c *fiber.Ctx)  error{
	email:=string(c.Request().Header.Peek("user_email"))
	data,err:=database.GetAllServicesCreated(email)
	if err!=nil {
		c.Status(409)
		return c.JSON(fiber.Map{"message":"Error Occured"})
	}else {
		return c.JSON(data)
	}
	
}

func GetAllSlotsOfService(c *fiber.Ctx)  error{
	email:=string(c.Request().Header.Peek("user_email"))
	data:=new(models.ServiceIdRequsest)
	if err := c.BodyParser(data); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":"Required arguments not found"})
	}
	slots,err:=database.GetAllSlotsOfService(email,data.ServiceID)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":err.Error()})
	}
	return c.JSON(slots)
}


func DeleteService(c *fiber.Ctx)  error{
	email:=string(c.Request().Header.Peek("user_email"))
	data:=new(models.ServiceIdRequsest)
	if err := c.BodyParser(data); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":"Required arguments not found"})
	}
	if database.DeleteService(email,data.ServiceID) {
		return c.JSON(fiber.Map{"message":"Done"})
	}else {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":"Delete failed"})
	}
}


func ApproveClientRequest(c *fiber.Ctx)  error{
	email:=string(c.Request().Header.Peek("user_email"))
	data:=new(models.UpdateRequest)
	if err := c.BodyParser(data); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":"Required arguments not found"})
	}
	if database.UpdateApproved(email,data.ServiceID,data.SlotId,data.Date) {
		return c.JSON(fiber.Map{"message":"Done"})
	}else{
		return c.JSON(fiber.Map{"message":"Delete failed"})
	}
}

func RemoveClientRequest(c *fiber.Ctx)  error{
	email:=string(c.Request().Header.Peek("user_email"))
	data:=new(models.UpdateRequest)
	if err := c.BodyParser(data); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":"Required arguments not found"})
	}
	if database.UpdateRemovedClient(email,data.ServiceID,data.SlotId,data.Date) {
		return c.JSON(fiber.Map{"message":"Done"})
	}else{
		return c.JSON(fiber.Map{"message":"Delete failed"})
	}
}