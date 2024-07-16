package routes

import (
	"scylla/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRoutesV1(
	dmsHandler *handler.DmsHandler,
	customerHandler *handler.CustomerHandler,
) *fiber.App {
	app := fiber.New()
	routes := app.Group("")
	//api third party
	routes.Get("/vehicles", dmsHandler.GetVehicle)

	//customer
	customerRouter := routes.Group("/customers")
	customerRouter.Get("", customerHandler.FindAllPaging)
	customerRouter.Get("/export", customerHandler.Export)
	customerRouter.Get("/:customerId", customerHandler.FindById)
	customerRouter.Post("", customerHandler.Create)
	customerRouter.Post("/batch", customerHandler.CreateBatch)
	customerRouter.Patch("/:customerId", customerHandler.Update)
	customerRouter.Delete("/batch", customerHandler.DeleteBatch)
	customerRouter.Post("/import", customerHandler.Import)

	return app
}
