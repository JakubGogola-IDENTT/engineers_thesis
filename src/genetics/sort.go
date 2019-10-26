package genetics

import (
	"fmt"
	"sort"
)

// byScore is wrapping type for sorting algorithm
type byScore []Specimen

// Len returns length of speciments array
func (s byScore) Len() int {
	return len(s)
}

// Swap swaps indices of two speciments in array
func (s byScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less returns speciment with smaller score
func (s byScore) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

// sortSpeciments sorts speciments in specified order (default - increasing)
// Returns sorted slice
func sortSpeciments(specimens []Specimen, reverse bool) []Specimen {
	sort.Sort(byScore(specimens))

	// Reverse if specified
	if reverse {
		sort.Sort(sort.Reverse(byScore(specimens)))
	}

	// printSpecimens(specimens)

	return specimens
}

func printSpecimens(specs []Specimen) {
	fmt.Println("*************")
	for _, spec := range specs {
		fmt.Println(spec.Score)
	}
}
