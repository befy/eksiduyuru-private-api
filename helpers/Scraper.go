package helpers

import (
	"eksiduyuru-private-api/models"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const BaseURL = "https://www.eksiduyuru.com"

type IScraper interface {
	Scrape() (*goquery.Document, error)
	GetPreviewPosts(doc *goquery.Document)
}

type Scraper struct {
}

func (s *Scraper) Scrape() (*goquery.Document, error) {
	res, err := http.Get(BaseURL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("burası var mı")
	return doc, nil
}

func (s *Scraper) GetPreviewPosts(doc *goquery.Document) *[]models.PostPreview {
	posts := make([]models.PostPreview, 0)
	var err error

	if err = generatePostList(doc, 1, &posts); err != nil {
		return nil
	}
	if err = generatePostList(doc, 0, &posts); err != nil {
		return nil
	}

	return &posts
}

func generatePostList(doc *goquery.Document, postType int, posts *[]models.PostPreview) error {
	mainSelector := fmt.Sprintf(".content .entry%d", postType)
	var err error
	doc.Find(mainSelector).Each(func(i int, s *goquery.Selection) {

		title := s.Find("h2.title.closed").Text()
		subtitle := s.Find(".bottomleft").Text()
		id := s.Find("ul > li:nth-child(1) > a:nth-child(2)").Text()

		if len(title) == 0 || len(subtitle) == 0 || len(id) == 0 {
			err = errors.New("field_missing")
			return
		}

		post := models.PostPreview{
			Title:    title,
			Subtitle: subtitle,
			ID:       id,
			Type:     postType,
		}
		*posts = append(*posts, post)
	})
	return err
}
