package main

import (
	"github.com/Proutyeahs/AZ-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowOrigins: "http://localhost:3000, http://localhost:8081",
	}))

	setupRoutes(app)

	app.Listen(":3000")
}
