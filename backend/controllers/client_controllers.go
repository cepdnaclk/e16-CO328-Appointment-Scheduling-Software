package controllers

import (
	"appoiment-backend/database"
	"appoiment-backend/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetClientServicses(c *fiber.Ctx)  error{
	data,err:=database.GetClientServicses();
	if err!=nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message":"Server Error"})
	}
	return c.JSON(data)
}

func GetAllSlotsOfClientService(c *fiber.Ctx)  error{
	req:=new(models.ClientServiceSlotRequest)
	if err := c.BodyParser(req); err != nil{
		
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":"Required arguments not found"})
	}
	data,err:=database.GetAllSlotsOfService(req.OwnerEmail,req.ServiceID)
	if err!=nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":err.Error()})
	}
	var res []*models.ClientDayDetailResponse
	for _, v := range *data{
		dateInst, err := time.Parse("2006-January-02", v.Date)
		if err != nil {
			fmt.Println(err)
		}
		if dateInst.Before(time.Now()) {
			continue
		}
		instday:=new(models.ClientDayDetailResponse)
		instday.Date=v.Date
		for _, w := range v.SlotList {
			instday.SlotList=append(instday.SlotList,models.ClientTimeSlot{
				ClientRequested:w.ClientRequested,
				SlotId:w.SlotId,
				Time:w.Time,
			})
		}
		res=append(res,instday)
	}
	return c.JSON(res)
}


func RequestingService(c *fiber.Ctx)  error{
	req:=new(models.ClientRequestingService)
	clientEmail:=string(c.Request().Header.Peek("user_email"))
	if err := c.BodyParser(req); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":"Required arguments not found"})
	}
	status,err:=database.ClientRequestingService(req.ServiceOwnerEmail,clientEmail,req.ServiceID,req.SlotId,req.Date)
	if err != nil {
		c.Status(409)
		return c.JSON(fiber.Map{"message":err})
	}
	return c.JSON(fiber.Map{"message":status})
}

func CancelRequestedService(c *fiber.Ctx)  error{
	req:=new(models.ClientRequestingService)
	clientEmail:=string(c.Request().Header.Peek("user_email"))
	if err := c.BodyParser(req); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":"Required arguments not found"})
	}
	
	if err:=database.CancelRequestedService(req.ServiceOwnerEmail,clientEmail,req.ServiceID,req.SlotId,req.Date); err != nil {
		c.Status(409)
		return c.JSON(fiber.Map{"message":err})
	}
	return c.JSON(fiber.Map{"message":"Done"})
}

func GetAllRequestedServices(c *fiber.Ctx)  error{
	clientEmail:=string(c.Request().Header.Peek("user_email"))
	data,err:=database.GetAllRequestedServices(clientEmail)
	if err != nil {
		c.Status(409)
		return c.JSON(fiber.Map{"message":err})
	}
	return c.JSON(data)
}

func SearchServices(c *fiber.Ctx)  error{
	req:=new(models.SearchServiceName)
	data,err:=database.SearchByName(req.ServiceName)
	if err != nil {
		c.Status(409)
		return c.JSON(fiber.Map{"message":err})
	}
	return c.JSON(data)
}

func SericeFindById(c *fiber.Ctx)  error{
	return nil
}