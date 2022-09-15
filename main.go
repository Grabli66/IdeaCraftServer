package main

import (
	"log"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")
	//port := "8080"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + port)
}
