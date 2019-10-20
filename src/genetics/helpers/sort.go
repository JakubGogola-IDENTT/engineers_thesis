package helpers

import (
	"sort"
	"thesis/genetics"
)

// ByScore is wrapping type for sorting algorithm
type ByScore []genetics.Speciment

// Len returns length of speciments array
func (s ByScore) Len() int {
	return len(s)
}

// Swap swaps indices of two speciments in array
func (s ByScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less returns speciment with smaller score
func (s ByScore) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

// SortSpeciments sorts speciments in specified order (default - increasing)
// Returns sorted slice
func SortSpeciments(speciments []genetics.Speciment, reverse bool) []genetics.Speciment {
	sort.Sort(ByScore(speciments))

	// Reverse if specified
	if reverse {
		sort.Reverse(ByScore(speciments))
	}

	return speciments
}
