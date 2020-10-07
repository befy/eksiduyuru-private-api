package main

import (
	"eksiduyuru-private-api/controllers"
	"eksiduyuru-private-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", middlewares.Scrape, controllers.GetPostList)

	app.Listen(":3000")
}
