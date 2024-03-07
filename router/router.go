package router

import (
	"fiber-boilerplate/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(noteController *controller.NoteController) *fiber.App {
	service := fiber.New()

	noteRouter := service.Group("/notes")
	noteRouter.Get("", noteController.FindAll)
	noteRouter.Get("/:noteId", noteController.FindById)
	noteRouter.Post("", noteController.Create)
	noteRouter.Delete("/:noteId", noteController.Delete)

	return service
}
