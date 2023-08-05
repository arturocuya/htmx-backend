package main

import (
	"text/template"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	type InitialData struct {
		Name string
	}

	app.Get("/", func(context *fiber.Ctx) error {
		tmpl, err := template.ParseFiles("./templates/index.html")

		if err != nil {
			return context.SendStatus(500)
		}

		data := InitialData{Name: "ayumi"}

		err = tmpl.Execute(context.Response().BodyWriter(), data)
		if err != nil {
			return context.SendStatus(500)
		}

		context.Response().Header.Set("Content-Type", "text/html")

		return nil
	})

	app.Listen(":8080")
}
