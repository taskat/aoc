package slices

import (
	"cmp"
	"fmt"
	"slices"
	"sort"

	"github.com/taskat/aoc/pkg/utils/iterutils"
	"github.com/taskat/aoc/pkg/utils/types"
)

// All returns true if all elements in the slice satisfy the predicate and false
// otherwise. It returns false at the first element that does not satisfy the
// predicate, and does not check the rest of the elements
func All[T any](slice []T, predicate func(T) bool) bool {
	return iterutils.All(slices.Values(slice), predicate)
}

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
func Copy[S types.Slice[T], T any](slice S) S {
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

// Equal returns true if the two slices are equal and false otherwise
func Equal[T comparable](slice1 []T, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
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
func First[S types.Slice[T], T any](slice S) T {
	if IsEmpty(slice) {
		panic("empty slice")
	}
	return slice[0]
}

// For applies the function f i times, and returns a slice with the results
func For[T any](i int, f func(int) T) []T {
	result := make([]T, i)
	for j := 0; j < i; j++ {
		result[j] = f(j)
	}
	return result
}

// ForEach applies the function f to each element of the slice
func ForEach[T any](slice []T, f func(T)) {
	for _, v := range slice {
		f(v)
	}
}

// ForEach_m applies the function f to each element of the slice. The function f
// receives a pointer to the element as an argument
func ForEach_m[T any](slice []T, f func(*T)) {
	for i := range slice {
		f(&slice[i])
	}
}

// IsEmpty returns true if the slice is empty and false otherwise
func IsEmpty[S types.Slice[T], T any](slice S) bool {
	return len(slice) == 0
}

// IsInBounds returns true if the index is within the bounds of the slice and
// false otherwise
func IsInBounds[T any](slice []T, index int) bool {
	return index >= 0 && index < len(slice)
}

// Last returns the last element of the slice. If the slice is empty, it panics
func Last[S types.Slice[T], T any](slice S) T {
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

// Max returns the maximum element in the slice. If the slice is empty, it panics
func Max[S types.Slice[T], T cmp.Ordered](slice S) T {
	max, _ := Max_i(slice)
	return max
}

// Max_i returns the maximum element in the slice. If the slice is empty, it panics
func Max_i[S types.Slice[T], T cmp.Ordered](slice S) (T, int) {
	if IsEmpty(slice) {
		panic("empty slice")
	}
	max := slice[0]
	index := 0
	for i, v := range slice {
		if v > max {
			max = v
			index = i
		}
	}
	return max, index
}

// Middle returns the middle element of the slice. If the slice has an even
// number of elements, it returns the first element of the second half. If the
// slice is empty, it panics
func Middle[S types.Slice[T], T any](slice S) T {
	if IsEmpty(slice) {
		panic("empty slice")
	}
	return slice[len(slice)/2]
}

// Min returns the minimum element in the slice. If the slice is empty, it panics
func Min[S types.Slice[T], T cmp.Ordered](slice S) T {
	min, _ := Min_i(slice)
	return min
}

// Min_i returns the minimum element in the slice. If the slice is empty, it panics
func Min_i[S types.Slice[T], T cmp.Ordered](slice S) (T, int) {
	if IsEmpty(slice) {
		panic("empty slice")
	}
	min := slice[0]
	index := 0
	for i, v := range slice {
		if v < min {
			min = v
			index = i
		}
	}
	return min, index
}

// Product returns the product of all the elements in the slice
func Product[S types.Slice[T], T types.Number](slice S) T {
	var product T = 1
	for _, v := range slice {
		product *= v
	}
	return product
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
		panic(fmt.Sprintf("index %d out of bounds: %d", index, len(slice)))
	}
	result := make([]T, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	return append(result, slice[index+1:]...)
}

// Repeat returns a new slice with the element repeated n times
func Repeat[T any](element T, n int) []T {
	result := make([]T, n)
	for i := range result {
		result[i] = element
	}
	return result
}

// Sort sorts the slice in increasing order
func Sort[S types.Slice[T], T cmp.Ordered](slice S, less func(T, T) bool) {
	sort.Slice(slice, func(i, j int) bool { return less(slice[i], slice[j]) })
}

// Sum returns the sum of all the elements in the slice
func Sum[S types.Slice[T], T types.Summable](slice S) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Split splits the slice into smaller slices, based on the given predicate
func Split[T any](slice []T, predicate func(T) bool) [][]T {
	result := make([][]T, 0)
	start := 0
	for i, v := range slice {
		if predicate(v) {
			result = append(result, slice[start:i])
			start = i + 1
		}
	}
	if start < len(slice) {
		result = append(result, slice[start:])
	}
	return result
}

// Swap swaps the elements at the given indices in the slice. If the indices are
// out of bounds, it panics
func Swap[T any](slice []T, i, j int) {
	if !IsInBounds(slice, i) {
		panic(fmt.Sprintf("index %d out of bounds: %d", i, len(slice)))
	}
	if !IsInBounds(slice, j) {
		panic(fmt.Sprintf("index %d out of bounds: %d", j, len(slice)))
	}
	slice[i], slice[j] = slice[j], slice[i]
}

// ToMap converts two slices into a map. The first slice contains the keys and
// the second slice contains the values. If the slices have different lengths,
// it panics
func ToMap[K comparable, V any](keys []K, values []V) map[K]V {
	if len(keys) != len(values) {
		panic("slices have different lengths")
	}
	result := make(map[K]V)
	for i, key := range keys {
		result[key] = values[i]
	}
	return result
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
