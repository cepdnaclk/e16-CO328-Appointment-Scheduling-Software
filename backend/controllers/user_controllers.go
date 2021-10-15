package controllers

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx)  error{
	return c.SendString("Hello, World!")
}

func SignUp(c *fiber.Ctx)  error{
	return c.SendString("Hello, World!")
}