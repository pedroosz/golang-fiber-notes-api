package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pedroosz/golang-fiber-notes-api/config"
	"github.com/pedroosz/golang-fiber-notes-api/routes"
)

func main() {
	app := fiber.New()

	config.Connect()

	routes.Setup(app)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(200)
	})

	log.Fatal(app.Listen(":3000"))
}
