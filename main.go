package main

import "fmt"

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
	ads := ScrapeAds()
	newAds := filter(ads, func(ad Ad) bool {
		return true
	})

	// TODO Actually only send new ads
	// Ads seem to be ordered in non-id order? Probably publish date.
	fmt.Printf("Posting %d new ads to Slack\n", len(newAds))

	for _, ad := range ads {
		slackError := SendAdToSlack(ad)
		if slackError != nil {
			fmt.Printf("Sending to Slack failed: %s\n", slackError)
		}
	}
}
