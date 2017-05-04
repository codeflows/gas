package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

type item struct {
	id          string
	category    string
	title       string
	url         string
	description string
	price       string
}

var idFromURL, _ = regexp.Compile("/(\\d+)$")

func ScrapeMuusikoidenNet() {
	doc, err := goquery.NewDocument("https://muusikoiden.net/tori/?category=0")
	if err != nil {
		log.Fatal(err)
	}

	elements := doc.Find("td.tori_title")

	items := make([]item, elements.Length())

	elements.Each(func(i int, titleContainer *goquery.Selection) {
		category := titleContainer.Find("b").Text()
		link := titleContainer.Find("a")
		title := link.Text()
		url, _ := link.Attr("href")
		id := idFromURL.FindStringSubmatch(url)[1]

		descriptionContainer := titleContainer.Parent().Siblings().Find("font.msg")
		description := descriptionContainer.Text()
		price := descriptionContainer.SiblingsFiltered("p").Text()

		item := item{id, category, title, url, description, price}
		items[i] = item
	})

	for _, item := range items {
		fmt.Printf("%s %s\n", item.title, item.price)
	}
}

func main() {
	ScrapeMuusikoidenNet()
}
