package router

import (
	"github.com/gofiber/fiber/v2"
	"restful-api-gorm-fiber/controller"
)

func NewRouter(noteController *controller.NoteController) *fiber.App {
	router := fiber.New()

	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to golang, fiber and GORM",
		})
	})

	router.Route("/note", func(router fiber.Router) {
		router.Post("/", noteController.Create)
		router.Get("", noteController.FindAll)
	})

	router.Route("/notes/:noteid", func(router fiber.Router) {
		router.Delete("", noteController.Delete)
		router.Get("", noteController.FindById)
		router.Patch("", noteController.Update)
	})
	return router
}
