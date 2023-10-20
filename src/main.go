package main

import (
	controller "example.com/gth-stack/src/modules"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    controller.IndexRouter(app)

    app.Listen(":42069")
}

func helper_func() {

}
