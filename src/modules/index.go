package controller

import (
	views "example.com/gth-stack/src/views"
	"github.com/gofiber/fiber/v2"
)

func IndexRouter(app *fiber.App) {
	app.Get("/", getLandingPage)
}

func getLandingPage(c *fiber.Ctx) error {
	// @todo: middleware so we just call this on context or something
	c.Response().Header.Add("Content-Type", "text/html")

	return views.Index().Render(c.Context(), c.Response().BodyWriter())
}

