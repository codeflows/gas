package main

// Ad is an advertisement in muusikoiden.net
type Ad struct {
	id          string
	category    string
	title       string
	url         string
	description string
	price       string
}

// Filter returns the array filtered by the predicate
func Filter(vs []Ad, predicate func(Ad) bool) []Ad {
	var filtered []Ad
	for _, v := range vs {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
