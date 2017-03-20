package deckMaker

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const (
	// URL is the target to scrape word definitions from
	URL = "http://ejje.weblio.jp/content/"
	// PATTERN is used by goquery for parsing HTMLs
	PATTERN = ".level0"
)

func getMeaning(word string) (string, error) {
	doc, err := goquery.NewDocument(URL + word)
	if err != nil {
		return "", fmt.Errorf("Error occured fetching a page. error = %v", err)
	}

	var definition string
	doc.Find(PATTERN).Each(func(i int, s *goquery.Selection) {
		definition += s.Text() + "<br>"
	})
	return definition, nil
}
