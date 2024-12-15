package set

import "github.com/taskat/aoc/pkg/utils/maps"

// Set is a set. It is implemented as a map with empty structs.
type Set[T comparable] map[T]struct{}

// FromSlice creates a new set from a slice.
func FromSlice[T comparable](s []T) Set[T] {
	result := make(Set[T])
	for _, e := range s {
		result[e] = struct{}{}
	}
	return result
}

// New creates a new set.
func New[T comparable]() Set[T] {
	return make(Set[T])
}

// Add adds an element to the set.
func (s Set[T]) Add(e T) {
	s[e] = struct{}{}
}

// Contains checks if the set contains an element.
func (s Set[T]) Contains(e T) bool {
	_, ok := s[e]
	return ok
}

// Delete deletes an element from the set.
func (s Set[T]) Delete(e T) {
	delete(s, e)
}

// Map applies a function to all elements of the set.
func (s Set[T]) Map(f func(T) T) Set[T] {
	fWrapper := func(e T, _ struct{}) (T, struct{}) {
		return f(e), struct{}{}
	}
	return maps.Map(s, fWrapper)
}

// Merge merges two sets into a new set.
func (s Set[T]) Merge(other Set[T]) Set[T] {
	result := make(Set[T])
	for k := range s {
		result.Add(k)
	}
	for k := range other {
		result.Add(k)
	}
	return result
}

// ToSlice converts the set to a slice.
func (s Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s))
	for k := range s {
		result = append(result, k)
	}
	return result
}
