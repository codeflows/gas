package main

import "fmt"

func main() {
	ads := ScrapeAds()
	for _, item := range ads {
		fmt.Printf("%s %s\n", item.title, item.price)
	}
}
