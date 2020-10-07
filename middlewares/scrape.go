package middlewares

import (
	"eksiduyuru-private-api/helpers"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
)

func Scrape(c *fiber.Ctx) error {
	var scraper helpers.Scraper
	var doc *goquery.Document
	var err error

	doc, err = scraper.Scrape()

	if err != nil {
		c.JSON(fiber.Map{
			"status":  "fail",
			"message": "deneme",
		})
		return err
	}

	c.Locals("scraper", scraper)
	c.Locals("doc", doc)
	c.Next()
	return nil
}
