package controllers

import (
	"appoiment-backend/models"

	"github.com/gofiber/fiber/v2"
)

func AddNewService(c *fiber.Ctx)  error{
	service:=new(models.AddNewServiceRequest)
	if err := c.BodyParser(service); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":"Required arguments not found"})
	}
	return nil
}


func GetAllServices(c *fiber.Ctx)  error{
	return nil
}

func GetAllSlotsOfService(c *fiber.Ctx)  error{
	return nil
}


func DeleteService(c *fiber.Ctx)  error{
	return nil
}


func ApproveClientRequest(c *fiber.Ctx)  error{
	return nil
}

func RemoveClientRequest(c *fiber.Ctx)  error{
	return nil
}