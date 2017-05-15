package main

import "fmt"

func main() {
	seenAdIds, readError := ReadSeenAdIds()
	if readError == nil {
		fmt.Printf("Read a list of %d already seen ad ids\n", len(seenAdIds))
	} else {
		fmt.Printf("Couldn't read already seen ad ids, falling back to an empty list. Error: %s\n", readError)
		seenAdIds = make(SeenAds)
	}

	ads := ScrapeAds()
	fmt.Printf("Read %d ads from Muusikoiden.net\n", len(ads))

	newAds := Filter(ads, func(ad Ad) bool {
		alreadySeen := seenAdIds[ad.id]
		return !alreadySeen
	})
	fmt.Printf("%d ads are new\n", len(newAds))

	return

	for _, ad := range ads {
		slackError := SendAdToSlack(ad)
		if slackError != nil {
			fmt.Printf("Sending to Slack failed: %s\n", slackError)
		}
	}
}
