package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) error {
	router := app.Group("/api")

	SetupNoteRoutes(router)

	return nil
}
