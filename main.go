package main

import "fmt"

func main() {
	fmt.Printf("Starting to read already seen ad ids")

	seenAdIds, readError := ReadSeenAdIds()
	if readError != nil {
		fmt.Printf("Couldn't read already seen ad ids, falling back to an empty list. Error: %s\n", readError)
		seenAdIds = make(SeenAds)
	} else if seenAdIds == nil {
		seenAdIds = make(SeenAds)
	}
	fmt.Printf("We've previously seen %d different ads\n", len(seenAdIds))

	ads := ScrapeAds()
	fmt.Printf("Read %d latest ads from Muusikoiden.net\n", len(ads))

	newAds := Filter(ads, func(ad Ad) bool {
		alreadySeen := seenAdIds[ad.id]
		return !alreadySeen
	})
	fmt.Printf("%d ads are new\n", len(newAds))

	for _, ad := range newAds {
		seenAdIds[ad.id] = true
		fmt.Printf("Sending Slack notification for ad %s\n", ad.id)
		slackError := SendAdToSlack(ad)
		if slackError != nil {
			fmt.Printf("Sending to Slack failed: %s\n", slackError)
		}
	}

	updateError := WriteSeenAdIds(seenAdIds)
	if updateError != nil {
		fmt.Printf("Writing updated ad id list to database failed: %s\n", updateError)
	}
}
