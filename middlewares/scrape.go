package middlewares

import (
	"eksiduyuru-private-api/helpers"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
)

func Scrape(c *fiber.Ctx) error {
	var scraper helpers.Scraper
	var doc *goquery.Document
	var err error

	var id = c.Params("id")
	path := c.Path()

	ok := strings.Contains(path, "post")

	if ok && len(id) != 0 {
		doc, err = scraper.Scrape(helpers.Post, id)
	} else {
		doc, err = scraper.Scrape(helpers.Home, "")
	}

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	c.Locals("scraper", scraper)
	c.Locals("doc", doc)
	c.Next()
	return nil
}
