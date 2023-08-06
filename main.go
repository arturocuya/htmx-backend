package main

import (
	"text/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

var indexTemplate *template.Template
var mainInputTemplate *template.Template

func init() {
	indexTemplate = template.Must(template.ParseFiles("./templates/index.html", "./templates/main-input.html"))
	mainInputTemplate = template.Must(template.ParseFiles("./templates/main-input.html"))
}

func main() {
	app := fiber.New()

	type InitialData struct {
		Name string
	}

	var messages []string

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.Next()
	})

	app.Get("/", func(context *fiber.Ctx) error {
		data := InitialData{Name: "ayumi"}

		err := indexTemplate.Execute(context.Response().BodyWriter(), data)
		if err != nil {
			return context.SendStatus(500)
		}

		return nil
	})

	app.Post("/send-message", func(context *fiber.Ctx) error {
		newMessage := utils.CopyString(context.FormValue("message"))
		messages = append(messages, newMessage)

		tmpl, _ := mainInputTemplate.New("").Parse(`{{range .}}<p>{{.}}</p>{{end}}{{template "main-input" .}}`)
		tmpl.Execute(context.Response().BodyWriter(), messages)

		return nil
	})

	app.Listen(":8080")
}
