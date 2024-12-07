package slices

import "github.com/taskat/aoc/pkg/utils/types"

// Any returns true if at least one element in the slice satisfies the predicate
// and false otherwise. It returns true at the first element that satisfies the
// predicate, and does not check the rest of the elements
func Any[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

// Contains returns true if the item is in the slice and false otherwise
func Contains[T comparable](slice []T, item T) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

// Copy returns a new slice with the same elements as the original
func Copy[S Slice[T], T any](slice S) S {
	result := make(S, len(slice))
	copy(result, slice)
	return result
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

// FindIndex returns the index of the first element that satisfies the predicate
// and -1 if no element satisfies the predicate
func FindIndex[T any](slice []T, predicate func(T) bool) int {
	for i, v := range slice {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// First returns the first element of the slice. If the slice is empty, it panics
func First[S Slice[T], T any](slice S) T {
	if IsEmpty(slice) {
		panic("empty slice")
	}
	return slice[0]
}

// ForEach applies the function f to each element of the slice
func ForEach[T any](slice []T, f func(T)) {
	for _, v := range slice {
		f(v)
	}
}

// IsEmpty returns true if the slice is empty and false otherwise
func IsEmpty[S Slice[T], T any](slice S) bool {
	return len(slice) == 0
}

// IsInBounds returns true if the index is within the bounds of the slice and
// false otherwise
func IsInBounds[T any](slice []T, index int) bool {
	return index >= 0 && index < len(slice)
}

// Last returns the last element of the slice. If the slice is empty, it panics
func Last[S Slice[T], T any](slice S) T {
	if IsEmpty(slice) {
		panic("empty slice")
	}
	return slice[len(slice)-1]
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

// Middle returns the middle element of the slice. If the slice has an even
// number of elements, it returns the first element of the second half. If the
// slice is empty, it panics
func Middle[S Slice[T], T any](slice S) T {
	if IsEmpty(slice) {
		panic("empty slice")
	}
	return slice[len(slice)/2]
}

// Reduce applies the function f to the elements of the slice and returns a
// single value. The function f receives the accumulated value and the current
// element as arguments
func Reduce[T any](slice []T, f func(T, T) T, initialValue T) T {
	functionWithIndex := func(accumulated T, current T, _ int) T { return f(accumulated, current) }
	return Reduce_i(slice, functionWithIndex, initialValue)
}

// Reduce_i applies the function f to the elements of the slice and returns a
// single value. The function f receives the index of the element as a third
// argument
func Reduce_i[T any](slice []T, f func(T, T, int) T, initialValue T) T {
	result := initialValue
	for i, value := range slice {
		result = f(result, value, i)
	}
	return result
}

// RemoveNth returns a new slice with the element at the given index removed
// If the index is out of bounds, it panics
func RemoveNth[T any](slice []T, index int) []T {
	if !IsInBounds(slice, index) {
		panic("index out of bounds")
	}
	result := make([]T, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	return append(result, slice[index+1:]...)
}

// Sum returns the sum of all the elements in the slice
func Sum[S Slice[T], T types.Summable](slice S) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Swap swaps the elements at the given indices in the slice. If the indices are
// out of bounds, it panics
func Swap[T any](slice []T, i, j int) {
	if !IsInBounds(slice, i) || !IsInBounds(slice, j) {
		panic("index out of bounds")
	}
	slice[i], slice[j] = slice[j], slice[i]
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
