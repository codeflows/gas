package main

import (
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// Ad is an advertisement in muusikoiden.net
type Ad struct {
	id          string
	category    string
	title       string
	url         string
	description string
	price       string
}

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
		category := titleContainer.Find("b").Text()
		link := titleContainer.Find("a")
		title := link.Text()
		url, _ := link.Attr("href")
		id := idFromURL.FindStringSubmatch(url)[1]

		descriptionContainer := titleContainer.Parent().Siblings().Find("font.msg")
		description := descriptionContainer.Text()
		price := descriptionContainer.SiblingsFiltered("p").Text()

		ad := Ad{id, category, title, url, description, price}
		ads[i] = ad
	})

	return ads
}
