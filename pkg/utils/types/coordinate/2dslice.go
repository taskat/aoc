package coordinate

import "github.com/taskat/aoc/pkg/utils/slices"

// Copy2DSlice creates a deep copy of the 2D slice
func Copy2DSlice[T any](slice [][]T) [][]T {
	result := make([][]T, len(slice))
	for i, row := range slice {
		result[i] = make([]T, len(row))
		copy(result[i], row)
	}
	return result
}

// CountIn2DSlice counts the number of elements in the 2D slice that satisfy the predicate
func CountIn2DSlice[T any](slice [][]T, predicate func(T) bool) int {
	count := 0
	for _, row := range slice {
		count += slices.Count(row, predicate)
	}
	return count
}

// Find2DSliceIndex returns the coordinates of the first element in the 2D slice
// that satisfies the predicate or (-1, -1) if no element satisfies the predicate
func Find2DSliceIndex[T any](slice [][]T, predicate func(T) bool) Coordinate2D[int] {
	for i, row := range slice {
		for j, elem := range row {
			if predicate(elem) {
				return FromIndexes(i, j)
			}
		}
	}
	return Coordinate2D[int]{X: -1, Y: -1}
}
