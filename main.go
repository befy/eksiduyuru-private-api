package main

import (
	"fmt"
	"log"
	"net/http"

	"eksiduyuru-private-api/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
)

const EksiDuyuruURL = "https://www.eksiduyuru.com"

type postPreview models.PostPreview

func Scrape() {
	res, err := http.Get(EksiDuyuruURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	posts := make([]models.PostPreview, 0)

	generateListPostsRequest(1, doc, &posts)
	generateListPostsRequest(0, doc, &posts)

	for _, post := range posts {
		fmt.Println(post, "\n----\n")
	}

}

func generateListPostsRequest(postType int, document *goquery.Document, posts *[]models.PostPreview) {
	mainSelector := fmt.Sprintf(".content .entry%d", postType)
	document.Find(mainSelector).Each(func(i int, s *goquery.Selection) {
		title := s.Find("h2.title.closed").Text()
		subtitle := s.Find(".bottomleft").Text()
		id := s.Find("ul > li:nth-child(1) > a:nth-child(2)").Text()
		post := models.PostPreview{
			Title:    title,
			Subtitle: subtitle,
			ID:       id,
			Type:     models.PostType{Type: 1},
		}
		*posts = append(*posts, post)
	})

}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
		})
	})

	app.Listen(":3000")
	Scrape()
}
