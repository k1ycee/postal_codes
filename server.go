package main

import (
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	jsonContent, err := ioutil.ReadFile("./postal_codes.json")

	if err != nil {
		log.Fatal(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(string(jsonContent))
	})

	app.Listen(":3000")
}
