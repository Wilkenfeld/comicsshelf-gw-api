package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Print("Starting server...")
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Print("Defaulting to port " + port)
	}
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
