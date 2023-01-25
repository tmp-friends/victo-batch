package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func requestHTML(url string) *goquery.Document {
	// Request the HTML page
	res, err := http.Get(url)
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

	return doc
}
