package service

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func ExtractTextFromURL(url string) (string, string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", "", err
	}

	title := doc.Find("title").Text()

	var content strings.Builder
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		content.WriteString(s.Text() + " ")
	})

	return title, content.String(), nil
}
