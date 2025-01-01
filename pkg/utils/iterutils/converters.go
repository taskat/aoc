package iterutils

import (
	"iter"

	"github.com/taskat/aoc/pkg/utils/types"
)

// NewFromFunc returns a new iterator that yields the values
// returned by the given function. The function takes no arguments.
// The length of the iterator is n
func NewFromFunc[T any](f func() T, n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < n; i++ {
			if !yield(f()) {
				return
			}
		}
	}
}

// NewFromFunc2 returns a new iterator that yields the key-value
// pairs returned by the given function. The function takes no
// arguments. The length of the iterator is n.
func NewFromFunc2[K comparable, V any](f func() (K, V), n int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i := 0; i < n; i++ {
			k, v := f()
			if !yield(k, v) {
				return
			}
		}
	}
}

// NewFromFuncIterations returns a new iterator that yields the
// values returned by the given function. The function takes the
// current iteration number as an argument. The length of the
// iterator is n.
func NewFromFuncIterations[T any](f func(int) T, n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < n; i++ {
			if !yield(f(i)) {
				return
			}
		}
	}
}

// NewFromFuncIterations2 returns a new iterator that yields the
// key-value pairs returned by the given function. The function
// takes the current iteration number as an argument. The length
// of the iterator is n.
func NewFromFuncIterations2[K comparable, V any](f func(int) (K, V), n int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i := 0; i < n; i++ {
			k, v := f(i)
			if !yield(k, v) {
				return
			}
		}
	}
}

// NewFromMap returns a new iterator that yields the key-value
// pairs of the given map.
func NewFromMap[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

// NewFromRepeat returns a new iterator that yields the same
// value a given number of times.
func NewFromRepeat[T any](value T, count int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < count; i++ {
			if !yield(value) {
				return
			}
		}
	}
}

// NewFromSlice returns a new iterator that yields the elements
// of the given slice.
func NewFromSlice[T any](slice []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range slice {
			if !yield(v) {
				return
			}
		}
	}
}

// NewFromSlice2 returns a new iterator that yields the key-value
// pairs of the given slice, where keys are the indices of the
// elements.
func NewFromSlice2[T any](slice []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range slice {
			if !yield(i, v) {
				return
			}
		}
	}
}

// ToMap collects the key-value pairs of the given iterator into a map.
// If the iterator is infinite, this function will not return.
// If the iterator is nil, this function will panic.
func ToMap[K comparable, V any](iter iter.Seq2[K, V]) map[K]V {
	panicIfNil[K, V](iter)
	m := make(map[K]V)
	iter(func(k K, v V) bool {
		m[k] = v
		return true
	})
	return m
}

// ToMapN collects the first n key-value pairs of the given iterator into a map.
// If n is greater than the length of the iterator, this function will return
// the whole iterator. If the iterator is nil, this function will panic.
func ToMapN[K comparable, V any](iter iter.Seq2[K, V], n uint) map[K]V {
	panicIfNil[K, V](iter)
	m := make(map[K]V, n)
	iter(func(k K, v V) bool {
		m[k] = v
		return uint(len(m)) < n
	})
	return m
}

// ToSlice collects the elements of the given iterator into a slice.
// If the iterator is infinite, this function will not return.
// If the iterator is nil, this function will panic.
func ToSlice[T any](iter iter.Seq[T]) []T {
	panicIfNil[any, T](iter)
	slice := make([]T, 0)
	iter(func(v T) bool {
		slice = append(slice, v)
		return true
	})
	return slice
}

// ToSlice2 collects the key-value pairs of the given iterator into a slice.
// The keys will be the indices of the elements.
// If the iterator is infinite, this function will not return.
// If the iterator is nil, this function will panic.
func ToSlice2[K types.Integer, V any](iter iter.Seq2[K, V]) []V {
	panicIfNil[K, V](iter)
	slice := make([]V, Len2(iter))
	iter(func(k K, v V) bool {
		slice[k] = v
		return true
	})
	return slice
}

// ToSliceN collects the first n elements of the given iterator into a slice.
// If n is greater than the length of the iterator, the function will return
// the whole iterator. If the iterator is nil, this function will panic.
func ToSliceN[T any](iter iter.Seq[T], n uint) []T {
	panicIfNil[any, T](iter)
	slice := make([]T, 0, n)
	iter(func(v T) bool {
		slice = append(slice, v)
		return uint(len(slice)) < n
	})
	return slice
}
