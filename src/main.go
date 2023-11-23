package main

import (
	controller "example.com/gth-stack/src/modules"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New(fiber.Config{
        AppName: "Venatus - Streaming Service",
    })

    controller.IndexRouter(app)

    app.Listen(":42069")
}

