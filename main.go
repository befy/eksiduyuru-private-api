package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const EKSIDUYURU_URL = "https://www.eksiduyuru.com"

func Scrape() {
	res, err := http.Get(EKSIDUYURU_URL)
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

	doc.Find(".content h2").Each(func(i int, s *goquery.Selection) {
		band := s.Text()
		fmt.Printf("Duyuru %d: %s\n\n", i, band)
	})
}

func main() {
	Scrape()
}
