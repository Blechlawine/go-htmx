package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: handlebars.New("./views", ".hbs"),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello World",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
