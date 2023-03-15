package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pedroosz/golang-fiber-notes-api/config"
	"github.com/pedroosz/golang-fiber-notes-api/models"
)

func GetNotes(ctx *fiber.Ctx) error {
	var notes []models.Note

	config.Database.Find(&notes)

	return ctx.Status(200).JSON(notes)
}

func AddNote(ctx *fiber.Ctx) error {
	note := new(models.Note)

	if err := ctx.BodyParser(note); err != nil {
		return ctx.Status(503).SendString(err.Error())
	}

	config.Database.Create(&note)

	return ctx.Status(200).JSON(note)
}

func DeleteNote(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var note models.Note

	result := config.Database.Delete(&note, id)

	if result.RowsAffected == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Not found",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Note deleted.",
	})
}

func GetNote(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var note models.Note

	result := config.Database.Find(&note, id)

	if result.RowsAffected == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Not found",
		})
	}

	return ctx.Status(200).JSON(note)
}
