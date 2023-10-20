package controller

import (
	view "example.com/gth-stack/src/views"
	"github.com/gofiber/fiber/v2"
)

func IndexRouter(app *fiber.App) {
    app.Get("/", getLandingPage)
}

func getLandingPage(c *fiber.Ctx) error {
    return view.Hello("Me").Render(c.Context(), c.Response().BodyWriter())
}

