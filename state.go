package main

// ReadSeenAdIds returns a list of ad ids we've already seen.
// Modified ads apparently resurface in the list,
// maybe we could also store a hash of the content?
func ReadSeenAdIds() map[string]bool {
	return make(map[string]bool)
}
