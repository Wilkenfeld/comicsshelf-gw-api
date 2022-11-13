package main

import (
	"context"
	"log"
	"os"

	_ "github.com/go-kivik/couchdb/v3"
	"github.com/go-kivik/kivik/v3"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Print("Starting server...")
	app := fiber.New()
	log.Print("Starting CouchDB driver...")
	client, err := kivik.New("couch", "https://comicsshelf-couchdb-jwayycdwga-oa.a.run.app/")
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.AllDBs(context.TODO())
	if err != nil {
		log.Print(err)
	}
	log.Print(res)
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
