package main

import (
	"encoding/json"
	"io/ioutil"
)

// SeenAds is the type used to represent seen ad ids
type SeenAds map[string]bool

const filename = "./seen_ad_ids"

// ReadSeenAdIds returns a map of ad ids we've already seen.
// Not an array since Go doesn't have `contains` for arrays out the box 😬 😎
func ReadSeenAdIds() (SeenAds, error) {
	file, fileError := ioutil.ReadFile(filename)

	if fileError != nil {
		return nil, fileError
	}

	var ids SeenAds
	jsonError := json.Unmarshal(file, &ids)

	if jsonError != nil {
		return nil, jsonError
	}
	return ids, nil
}

func WriteSeenAdIds(seenAds SeenAds) error {
	return nil
}
