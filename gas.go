package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeMuusikoidenNet() {
	doc, err := goquery.NewDocument("https://muusikoiden.net/tori/?category=0")
	if err != nil {
		log.Fatal(err)
	}

	idFromURL, _ := regexp.Compile("/(\\d+)$")

	doc.Find("td.tori_title").Each(func(i int, title *goquery.Selection) {
		// TODO "Myydään" ääkköset are mangled
		category := title.Find("b").Text()
		link := title.Find("a")
		name := link.Text()
		url, _ := link.Attr("href")
		id := idFromURL.FindStringSubmatch(url)[1]
		fmt.Printf("%s %s %s %s\n", category, name, url, id)
	})
}

func main() {
	ScrapeMuusikoidenNet()
}
