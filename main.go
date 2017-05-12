package main

import (
	"fmt"
	"io/ioutil"
)

func filter(vs []Ad, f func(Ad) bool) []Ad {
	vsf := make([]Ad, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func main() {
	lastSeenAdID, err := ioutil.ReadFile("last_seen")
	if err != nil {
		println("No last_seen state")
	}
	fmt.Printf("Last seen is %s\n", string(lastSeenAdID))

	ads := ScrapeAds()
	newAds := filter(ads, func(ad Ad) bool {
		return true
	})

	// Ads seem to be ordered in non-id order? Probably publish date.
	fmt.Printf("%d new ads since last run\n", len(newAds))

	for _, ad := range ads {
		fmt.Printf("%s %s %s\n", ad.id, ad.title, ad.price)
		slackError := SendAdToSlack(ad)
		if slackError != nil {
			fmt.Printf("Sending to Slack failed: %s\n", slackError)
		}
	}
}
