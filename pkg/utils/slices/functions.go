package slices

import "github.com/taskat/aoc/pkg/utils/types"

// Contains returns true if the item is in the slice and false otherwise
func Contains[T comparable](slice []T, item T) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

// Count returns the number of elements in the slice that satisfy the predicate
func Count[T any](slice []T, predicate func(T) bool) int {
	count := 0
	for _, v := range slice {
		if predicate(v) {
			count++
		}
	}
	return count
}

// Filter returns a new slice with the elements that satisfy the predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Find returns the first element that satisfies the predicate and a boolean
func Find[T any](arr []T, predicate func(T) bool) (T, bool) {
	for _, v := range arr {
		if predicate(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// Map applies the function f to each element of the slice and returns a
// new slice with the results
func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// Map_i applies the function f to each element of the slice and returns a
// new slice with the results. The function f receives the index of the element
// as a second argument
func Map_i[T, U any](slice []T, f func(T, int) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v, i)
	}
	return result
}

// Sum returns the sum of all the elements in the slice
func Sum[T types.Summable](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// ZipWith applies the function f to each pair of elements from the two slices
// and returns a new slice with the results. If the slices have different
// lengths, the result will have the length of the shortest slice
func ZipWith[T, U, V any](slice1 []T, slice2 []U, f func(T, U) V) []V {
	minLen := len(slice1)
	if len(slice2) < minLen {
		minLen = len(slice2)
	}
	result := make([]V, minLen)
	for i := 0; i < minLen; i++ {
		result[i] = f(slice1[i], slice2[i])
	}
	return result
}
