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

	ok := strings.Contains(path, "posts")

	if ok && len(id) != 0 {
		doc, err = scraper.Scrape(helpers.Post)
	} else {
		doc, err = scraper.Scrape(helpers.Home)
	}

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
