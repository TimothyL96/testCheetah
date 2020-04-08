package main

// Store recipient to Name, Id, Tags
// tagsMap will be generated later, which is used for faster searching of tags
type recipient struct {
	Name    string
	Id      int
	Tags    []string
	tagsMap map[string]struct{}
}

// Used to store recipient wrapper in the data
type recipientWrapper struct {
	JsonArrayData []recipient `json:"recipients"`
}

// Generate the tagsMap of the recipient with type <string, struct{}>
// We use struct{} here as value because Golang doesn't have Set, or map without value.
// When storing value below, we use struct{}{} or empty struct, as it actually doesn't need any memory to be stored
func (r *recipient) generateTagsMap() {
	// Create and preallocate the size of tagsMap with the already known size of Tags to prevent unnecessary new allocation
	r.tagsMap = make(map[string]struct{}, len(r.Tags))

	// Traverse the slice/list Tags to generate the map
	for _, t := range r.Tags {
		r.tagsMap[t] = struct{}{}
	}
}

// Compare if the recipient has at least 2 similar tags with the recipient from the argument input and returns a boolean value
func (r *recipient) hasTwoOrMoreSimilarTags(compareRecipient recipient) bool {
	count := 0

	for _, t := range r.Tags {
		// Tag 't' exists in tagsMap if ok is true
		if _, ok := compareRecipient.tagsMap[t]; ok {
			count++

			if count >= 2 {
				return true
			}
		}
	}

	return false
}
