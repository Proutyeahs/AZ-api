package main

import (
	"github.com/Proutyeahs/AZ-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.GetTiles)

	app.Post("/tile", handlers.CreateTile)
}
