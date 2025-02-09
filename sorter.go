package main

// SortSearchResults sorts search results.
func SortSearchResults(searchResults []SearchResult) []SearchResult {
	copied := make([]SearchResult, len(searchResults))
	// At most 10 items, so bubble sort is enough.
	copy(copied, searchResults)

	for i := 0; i < len(copied); i++ {
		for j := i + 1; j < len(copied); j++ {
			if copied[i].SearchCount < copied[j].SearchCount {
				copied[i], copied[j] = copied[j], copied[i]
			}
		}
	}

	return copied
}
