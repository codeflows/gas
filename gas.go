package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

type item struct {
	id       string
	category string
	title    string
	url      string
}

func ScrapeMuusikoidenNet() {
	doc, err := goquery.NewDocument("https://muusikoiden.net/tori/?category=0")
	if err != nil {
		log.Fatal(err)
	}

	idFromURL, _ := regexp.Compile("/(\\d+)$")

	doc.Find("td.tori_title").Each(func(i int, titleContainer *goquery.Selection) {
		// TODO "Myydään" ääkköset are mangled
		category := titleContainer.Find("b").Text()
		link := titleContainer.Find("a")
		title := link.Text()
		url, _ := link.Attr("href")
		id := idFromURL.FindStringSubmatch(url)[1]
		item := item{id, category, title, url}
		fmt.Printf("%s\n", item)
	})
}

func main() {
	ScrapeMuusikoidenNet()
}
