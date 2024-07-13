package router

import (
	"scylla/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(
	dmsHandler *handler.DmsHandler,
	customerHandler *handler.CustomerHandler,
) *fiber.App {
	service := fiber.New()

	//api third party
	service.Get("/vehicles", dmsHandler.GetVehicle)

	//customer
	customerRouter := service.Group("/customers")
	customerRouter.Get("", customerHandler.FindAllPaging)
	customerRouter.Get("/:customerId", customerHandler.FindById)
	customerRouter.Post("", customerHandler.Create)
	customerRouter.Post("/batch", customerHandler.CreateBatch)
	customerRouter.Patch("/:customerId", customerHandler.Update)
	customerRouter.Delete("/batch", customerHandler.DeleteBatch)

	return service
}
