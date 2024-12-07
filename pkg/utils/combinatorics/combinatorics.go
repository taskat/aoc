package combinatorics

import "github.com/taskat/aoc/pkg/utils/intutils"

// CartesianProduct returns the cartesian product of the element with itself length times
// The result is a slice of slices, where each slice is a possible combination of the elements
func CartesianProduct[T any](element []T, length int) [][]T {
	if length == 0 {
		return [][]T{{}}
	}
	result := make([][]T, 0, intutils.Power(len(element), length))
	for _, e := range element {
		rest := CartesianProduct(element, length-1)
		for _, r := range rest {
			result = append(result, append([]T{e}, r...))
		}
	}
	return result
}
