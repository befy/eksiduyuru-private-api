package controllers

import (
	"eksiduyuru-private-api/helpers"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
)

func GetPostList(c *fiber.Ctx) error {
	var doc *goquery.Document = c.Locals("doc").(*goquery.Document)
	var scraper helpers.Scraper = c.Locals("scraper").(helpers.Scraper)
	var err error

	posts := scraper.GetPreviewPosts(doc)

	if posts == nil {
		fmt.Println("error 2 var", err)
		c.JSON(fiber.Map{
			"status":  "failure",
			"message": "no_post_data",
		})
		return nil
	}
	c.JSON(fiber.Map{
		"status": "success",
		"data":   posts,
	})
	return nil
}

func GetDuyuruContent(c *fiber.Ctx) error {
	var doc *goquery.Document = c.Locals("doc").(*goquery.Document)
	var scraper helpers.Scraper = c.Locals("scraper").(helpers.Scraper)

	entries := scraper.GetEntryContent(doc)
	c.JSON(fiber.Map{
		"status": "success",
		"data":   entries,
	})
	return nil
}
