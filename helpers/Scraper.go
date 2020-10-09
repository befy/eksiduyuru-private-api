package helpers

import (
	"eksiduyuru-private-api/models"
	"eksiduyuru-private-api/utils"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const BaseURL = "https://www.eksiduyuru.com"

type Endpoint int

const (
	Home   Endpoint = iota // 0
	Post                   // 1
	Search                 // 2
)

type IScraper interface {
	Scrape(action Endpoint, id uint64) (*goquery.Document, error)
	GetPreviewPosts(doc *goquery.Document)
	GetEntryContent(doc *goquery.Document)
}

type Scraper struct{}

func (s *Scraper) Scrape(action Endpoint, id string) (*goquery.Document, error) {
	url := requestURLFactory(action, id)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		//log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		if res.StatusCode == 404 {
			return nil, errors.New("not_found_data")
		}
		return nil, errors.New("external_error")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return doc, nil
}

func (s *Scraper) GetPreviewPosts(doc *goquery.Document) *[]models.PostPreview {
	posts := make([]models.PostPreview, 0)
	var err error

	if err = generatePostList(doc, 1, &posts); err != nil {
		fmt.Println(err)
		return nil
	}
	if err = generatePostList(doc, 0, &posts); err != nil {
		return nil
	}

	return &posts
}

func (s *Scraper) GetEntryContent(doc *goquery.Document) *[]models.Entry {
	entries := make([]models.Entry, 0)
	getDuyuruContent(doc, &entries)

	return &entries
}

func generatePostList(doc *goquery.Document, postType int, posts *[]models.PostPreview) error {
	mainSelector := fmt.Sprintf(".content .entry%d", postType)
	var err error
	doc.Find(mainSelector).Each(func(i int, s *goquery.Selection) {

		title := s.Find("h2.title.closed").Text()
		subtitle := s.Find(".bottomleft").Text()
		strID := s.Find("ul > li:nth-child(1) > a:nth-child(2)").Text()

		id := utils.GetID(strID)

		if len(title) == 0 || len(subtitle) == 0 || id == 0 {
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

func getDuyuruContent(doc *goquery.Document, entries *[]models.Entry) {
	context := doc.Find(".content .entrybody").Text()
	var body string
	if len(context) == 0 {
		return
	}
	body = strings.TrimSpace(context)
	author := doc.Find(".entryhead > ul > li:nth-child(1) > a:nth-child(1)").Text()

	entry := models.Entry{
		Text:      body,
		Author:    author,
		CreatedAt: "a",
	}

	*entries = append(*entries, entry)

	doc.Find(".answerscontainer .answer").Each(func(i int, s *goquery.Selection) {
		answer := strings.TrimSpace(s.Find(".answerbody").Text())
		entryInfo := s.Find("ul.duans.poster > li").Text()

		info := utils.ParseAuthorInfo(entryInfo)

		if len(info) == 0 {
			return
		}

		entry := models.Entry{
			Text:      answer,
			Author:    info["author"],
			CreatedAt: info["date"],
		}

		*entries = append(*entries, entry)
	})

}

func requestURLFactory(reqType Endpoint, id string) string {
	switch reqType {
	case Home:
		return BaseURL
	case Post:
		url := BaseURL
		return url + "/duyuru/" + id
	default:
		return "a"
	}
}
