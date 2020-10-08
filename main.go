package main

import (
	"eksiduyuru-private-api/controllers"
	"eksiduyuru-private-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/post", middlewares.Scrape, controllers.GetPostList)
	app.Get("/post/:id", middlewares.Scrape, controllers.GetDuyuruContent)

	app.Listen(":3000")
}
