package main

import "fmt"

func main() {
	ads := ScrapeAds()
	fmt.Printf("Read %d ads from Muusikoiden.net\n", len(ads))

	seenAdIds := ReadSeenAdIds()

	newAds := Filter(ads, func(ad Ad) bool {
		alreadySeen := seenAdIds[ad.id]
		return !alreadySeen
	})
	fmt.Printf("Posting %d new ads to Slack\n", len(newAds))

	return

	for _, ad := range ads {
		slackError := SendAdToSlack(ad)
		if slackError != nil {
			fmt.Printf("Sending to Slack failed: %s\n", slackError)
		}
	}
}
