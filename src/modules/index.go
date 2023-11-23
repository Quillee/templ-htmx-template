package controller

import (
	views "example.com/gth-stack/src/views"
	"github.com/gofiber/fiber/v2"
)

func IndexRouter(app *fiber.App) {
	app.Get("/", get_landing_page)
    app.Get("/ui-docs", get_ui_doc_page)

}

func get_landing_page(c *fiber.Ctx) error {
	// @todo: middleware so we just call this on context or something
	c.Response().Header.Add("Content-Type", "text/html")

	return views.Index().Render(c.Context(), c.Response().BodyWriter())
}

func get_ui_doc_page(c *fiber.Ctx) error {
	c.Response().Header.Add("Content-Type", "text/html")
    return views.UiTest().Render(c.Context(), c.Response().BodyWriter())
}

