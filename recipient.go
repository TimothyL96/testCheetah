package main

type recipient struct {
	Name    string
	Id      int
	Tags    []string
	tagsMap map[string]struct{}
}

// For mapping JSON use
type recipientWrapper struct {
	JsonArrayData []recipient `json:"recipients"`
}

func (r *recipient) generateTagsMap() {
	r.tagsMap = make(map[string]struct{}, len(r.Tags))

	for _, t := range r.Tags {
		r.tagsMap[t] = struct{}{}
	}
}

func (r *recipient) hasTwoOrMoreSimilarTags(compareRecipient recipient) bool {
	count := 0

	for _, t := range r.Tags {
		if _, ok := compareRecipient.tagsMap[t]; ok {
			count++

			if count >= 2 {
				return true
			}
		}
	}

	return false
}
