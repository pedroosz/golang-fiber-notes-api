package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pedroosz/golang-fiber-notes-api/controllers"
)

func SetupNoteRoutes(router fiber.Router) {
	noteRouter := router.Group("/notes")

	noteRouter.Get("/", controllers.GetNotes)
	noteRouter.Post("/", controllers.AddNote)
	noteRouter.Get("/:id", controllers.GetNote)
	noteRouter.Delete("/:id", controllers.DeleteNote)
}
