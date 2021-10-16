package routes

import (
	"appoiment-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserSetup(app *fiber.App)  {
	app.Post("login",controllers.Login)
	app.Post("signup",controllers.SignUp)
}