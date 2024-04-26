package handlers

import (
	"github.com/Proutyeahs/AZ-api/database"
	"github.com/Proutyeahs/AZ-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetTiles(c *fiber.Ctx) error {
	tiles := []models.Tile{}
	
	database.DB.Db.Find(&tiles)

	return c.Status(200).JSON(tiles)
}

func CreateTile(c *fiber.Ctx) error {
	tile := new(models.Tile)
	if err := c.BodyParser(tile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&tile)

	return c.Status(200).JSON(tile)
}