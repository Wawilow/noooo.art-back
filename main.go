package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func handleStatus(c *fiber.Ctx) error {
	response := fiber.Map{
		"status": "success",
	}

	return c.JSON(response)
}

func main() {
	app := fiber.New()

	app.Get("/api/status", handleStatus)

	log.Println("Starting server on :3000...")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
