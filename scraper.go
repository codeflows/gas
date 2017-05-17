package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var idFromURL, _ = regexp.Compile("/(\\d+)$")

// ScrapeAds returns a list of ads from muusikoiden.net
func ScrapeAds() []Ad {
	doc, err := goquery.NewDocument("https://muusikoiden.net/tori/?category=0")
	if err != nil {
		log.Fatal(err)
	}

	elements := doc.Find("td.tori_title")

	ads := make([]Ad, elements.Length())

	elements.Each(func(i int, titleContainer *goquery.Selection) {
		category := strings.Replace(titleContainer.Find("b").Text(), ":", "", -1)
		link := titleContainer.Find("a")
		title := link.Text()
		path, _ := link.Attr("href")
		url := "https://muusikoiden.net" + path
		id := idFromURL.FindStringSubmatch(path)[1]

		descriptionContainer := titleContainer.Parent().Siblings().Find("font.msg")
		description := descriptionContainer.Text()
		price := descriptionContainer.SiblingsFiltered("p").Text()

		ad := Ad{id, category, title, url, description, price}
		ads[i] = ad
	})

	return ads
}
