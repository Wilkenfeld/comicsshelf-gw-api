package main;

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	log.Print("Starting server...")
	app := fiber.New()
	app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
	log.Fatal(app.Listen(os.Getenv("PORT")))
}