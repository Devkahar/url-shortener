package routes

import (
	"github.com/Devkahar/url-shortener/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")

	db := database.CreateClient(0)
	defer db.Close()
	value, err := db.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Short not found in database"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot connect to database"})
	}

	rInc := database.CreateClient(1)
	defer rInc.Close()

	_ = rInc.Incr(database.Ctx, "counter")
	return c.Redirect(value, fiber.StatusMovedPermanently)
}
