package routes

import (
	"appoiment-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ServiceSetup(app *fiber.App)  {
	app.Post("add-new-service",controllers.AddNewService)
	app.Get("get-all-created-services",controllers.GetAllServices)
	app.Post("get-slots-created-service",controllers.GetAllSlotsOfService)
	app.Post("delete-service",controllers.DeleteService)
	app.Post("approve-client-request",controllers.ApproveClientRequest)
	app.Post("remove-client-request",controllers.RemoveClientRequest)
}