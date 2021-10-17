package routes

import (
	"appoiment-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ClientSetup(app *fiber.App)  {
	app.Get("newest-services",controllers.GetClientServicses)
	app.Post("get-slots-client-service",controllers.GetAllSlotsOfClientService)
	app.Post("request-service",controllers.RequestingService)
	app.Post("cancel-requested-service",controllers.CancelRequestedService)
	app.Post("getall-requested-services",controllers.GetAllRequestedServices)
	app.Post("search-services",controllers.SearchServices)
	app.Post("service-find-by-Id",controllers.SericeFindById)
}