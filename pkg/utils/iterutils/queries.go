package iterutils

import (
	"cmp"
	"iter"

	"github.com/taskat/aoc/pkg/utils/types"
)

type (
	// Predicate is a function that takes a value and returns a boolean
	Predicate[T any] func(T) bool
	// Predicate2 is a function that takes a key and a value and returns a boolean
	Predicate2[K comparable, V any] func(K, V) bool
)

// All returns true if all elements in the iter satisfy the predicate and
// false otherwise. It returns false at the first element that does not
// satisfy the predicate, and does not check the rest of the elements.
// If the iter is empty, it returns true. If the iter is nil, it panics.
func All[T any](iter iter.Seq[T], predicate Predicate[T]) bool {
	panicIfNil[any, T](iter)
	for v := range iter {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// All2 returns true if all key-value pairs in the iter satisfy the predicate
// and false otherwise. It returns false at the first element that does not
// satisfy the predicate, and does not check the rest of the elements.
// If the iter is empty, it returns true. If the iter is nil, it panics.
func All2[K comparable, V any](iter iter.Seq2[K, V], predicate Predicate2[K, V]) bool {
	panicIfNil[K, V](iter)
	for k, v := range iter {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element in the iter satisfies the predicate
// and false otherwise. It returns true at the first element that satisfies the
// predicate, and does not check the rest of the elements.
// If the iter is empty, it returns false. If the iter is nil, it panics.
func Any[T any](iter iter.Seq[T], predicate Predicate[T]) bool {
	panicIfNil[any, T](iter)
	for v := range iter {
		if predicate(v) {
			return true
		}
	}
	return false
}

// Any2 returns true if at least one key-value pair in the iter satisfies the
// predicate and false otherwise. It returns true at the first element that
// satisfies the predicate, and does not check the rest of the elements.
// If the iter is empty, it returns false. If the iter is nil, it panics.
func Any2[K comparable, V any](iter iter.Seq2[K, V], predicate Predicate2[K, V]) bool {
	panicIfNil[K, V](iter)
	for k, v := range iter {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

// Contains returns true if the item is in the iter and false otherwise.
// If the iter is nil, it panics.
func Contains[T comparable](iter iter.Seq[T], item T) bool {
	panicIfNil[any, T](iter)
	for v := range iter {
		if v == item {
			return true
		}
	}
	return false
}

// Count returns the number of elements in the iter that satisfy the predicate.
// If the iter is nil, it panics.
func Count[T any](iter iter.Seq[T], predicate Predicate[T]) int {
	panicIfNil[any, T](iter)
	count := 0
	for v := range iter {
		if predicate(v) {
			count++
		}
	}
	return count
}

// Count2 returns the number of key-value pairs in the iter that satisfy the
// predicate. If the iter is nil, it panics.
func Count2[K comparable, V any](iter iter.Seq2[K, V], predicate Predicate2[K, V]) int {
	panicIfNil[K, V](iter)
	count := 0
	for k, v := range iter {
		if predicate(k, v) {
			count++
		}
	}
	return count
}

// Equal returns true if the values in the two iters are equal and false otherwise.
// If either one of the iters is nil, it panics.
func Equal[T comparable](iter1, iter2 iter.Seq[T]) bool {
	panicIfNil[any, T](iter1)
	panicIfNil[any, T](iter2)
	next1, stop1 := iter.Pull(iter1)
	defer stop1()
	next2, stop2 := iter.Pull(iter2)
	defer stop2()
	for {
		v1, ok1 := next1()
		v2, ok2 := next2()
		if !ok1 {
			return !ok2
		}
		if ok1 != ok2 || v1 != v2 {
			return false
		}
	}
}

// Equal2 returns true if the key-value pairs in the two iters are equal and false otherwise.
// If either one of the iters is nil, it panics.
func Equal2[K comparable, V comparable](iter1, iter2 iter.Seq2[K, V]) bool {
	panicIfNil[K, V](iter1)
	panicIfNil[K, V](iter2)
	next1, stop1 := iter.Pull2(iter1)
	defer stop1()
	next2, stop2 := iter.Pull2(iter2)
	defer stop2()
	for {
		k1, v1, ok1 := next1()
		k2, v2, ok2 := next2()
		if !ok1 {
			return !ok2
		}
		if ok1 != ok2 || k1 != k2 || v1 != v2 {
			return false
		}
	}
}

// Find returns the first element that satisfies the predicate and a boolean.
// that indicates if such an element was found. If the iter is nil, it panics.
func Find[T any](iter iter.Seq[T], predicate Predicate[T]) (T, bool) {
	panicIfNil[any, T](iter)
	for v := range iter {
		if predicate(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// Find2 returns the first key-value pair that satisfies the predicate and a boolean.
// that indicates if such a pair was found. If the iter is nil, it panics.
func Find2[K comparable, V any](iter iter.Seq2[K, V], predicate Predicate2[K, V]) (K, V, bool) {
	panicIfNil[K, V](iter)
	for k, v := range iter {
		if predicate(k, v) {
			return k, v, true
		}
	}
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// First returns the first element in the iter. If the iter is empty or nil, it panics.
func First[T any](iter iter.Seq[T]) T {
	panicIfNil[any, T](iter)
	if IsEmpty(iter) {
		panic(emptyIterPanicMsg)
	}
	for v := range iter {
		return v
	}
	panic("unreachable")
}

// First2 returns the first key-value pair in the iter. If the iter is empty or nil, it panics.
func First2[K comparable, V any](iter iter.Seq2[K, V]) (K, V) {
	panicIfNil[K, V](iter)
	if IsEmpty2(iter) {
		panic(emptyIterPanicMsg)
	}
	for k, v := range iter {
		return k, v
	}
	panic("unreachable")
}

// ForEach applies the function to each element in the iter. If the iter is nil, it panics.
func ForEach[T any](iter iter.Seq[T], f func(T)) {
	panicIfNil[any, T](iter)
	for v := range iter {
		f(v)
	}
}

// ForEach2 applies the function to each key-value pair in the iter. If the iter is nil, it panics.
func ForEach2[K comparable, V any](iter iter.Seq2[K, V], f func(K, V)) {
	panicIfNil[K, V](iter)
	for k, v := range iter {
		f(k, v)
	}
}

// IsEmpty returns true if the iter is empty and false otherwise. If the iter is nil, it panics.
func IsEmpty[T any](iter iter.Seq[T]) bool {
	panicIfNil[any, T](iter)
	return Len(iter) == 0
}

// IsEmpty2 returns true if the iter is empty and false otherwise. If the iter is nil, it panics.
func IsEmpty2[K comparable, V any](iter iter.Seq2[K, V]) bool {
	panicIfNil[K, V](iter)
	return Len2(iter) == 0
}

// IsValidKey returns true if the key is in the iter and false otherwise. If the iter is nil, it panics.
func IsValidKey[K comparable, V any](iter iter.Seq2[K, V], key K) bool {
	panicIfNil[K, V](iter)
	return Contains(Keys(iter), key)
}

// Last returns the last element in the iter. If the iter is empty or nil, it panics.
func Last[T any](iter iter.Seq[T]) T {
	panicIfNil[any, T](iter)
	if IsEmpty(iter) {
		panic(emptyIterPanicMsg)
	}
	var last T
	for last = range iter {
		// no-op
	}
	return last
}

// Last2 returns the last key-value pair in the iter. If the iter is empty or nil, it panics.
func Last2[K comparable, V any](iter iter.Seq2[K, V]) (K, V) {
	panicIfNil[K, V](iter)
	if IsEmpty2(iter) {
		panic(emptyIterPanicMsg)
	}
	var lastK K
	var lastV V
	for lastK, lastV = range iter {
		// no-op
	}
	return lastK, lastV
}

// Len returns the number of elements in the iter. If the iter is nil, it panics.
func Len[T any](iter iter.Seq[T]) int {
	panicIfNil[any, T](iter)
	count := 0
	for range iter {
		count++
	}
	return count
}

// Len2 returns the number of key-value pairs in the iter. If the iter is nil, it panics.
func Len2[K comparable, V any](iter iter.Seq2[K, V]) int {
	panicIfNil[K, V](iter)
	count := 0
	for range iter {
		count++
	}
	return count
}

// Max returns the maximum element in the iter. If the iter is empty or nil, it panics.
func Max[T cmp.Ordered](iter iter.Seq[T]) T {
	panicIfNil[any, T](iter)
	if IsEmpty(iter) {
		panic(emptyIterPanicMsg)
	}
	max := First(iter)
	for v := range iter {
		if v > max {
			max = v
		}
	}
	return max
}

// Max2 returns the maximum key-value pair in the iter. If the iter is empty or nil, it panics.
func Max2[K comparable, V cmp.Ordered](iter iter.Seq2[K, V]) (K, V) {
	panicIfNil[K, V](iter)
	if IsEmpty2(iter) {
		panic(emptyIterPanicMsg)
	}
	maxK, maxV := First2(iter)
	for k, v := range iter {
		if v > maxV {
			maxK, maxV = k, v
		}
	}
	return maxK, maxV
}

// Middle returns the middle element of the iter. If the iter has an even
// number of elements, it returns the first element of the second half.
// If the iter is empty or nil, it panics.
func Middle[T any](iter iter.Seq[T]) T {
	panicIfNil[any, T](iter)
	if IsEmpty(iter) {
		panic(emptyIterPanicMsg)
	}
	length := Len(iter)
	return First(Skip(iter, length/2))
}

// Middle2 returns the middle key-value pair of the iter. If the iter has an even
// number of elements, it returns the first key-value pair of the second half.
// If the iter is empty or nil, it panics.
func Middle2[K comparable, V any](iter iter.Seq2[K, V]) (K, V) {
	panicIfNil[K, V](iter)
	if IsEmpty2(iter) {
		panic(emptyIterPanicMsg)
	}
	length := Len2(iter)
	return First2(Skip2(iter, length/2))
}

// Min returns the minimum element in the iter. If the iter is empty or nil, it panics.
func Min[T cmp.Ordered](iter iter.Seq[T]) T {
	panicIfNil[any, T](iter)
	if IsEmpty(iter) {
		panic(emptyIterPanicMsg)
	}
	min := First(iter)
	for v := range iter {
		if v < min {
			min = v
		}
	}
	return min
}

// Min2 returns the minimum key-value pair in the iter. If the iter is empty or nil, it panics.
func Min2[K comparable, V cmp.Ordered](iter iter.Seq2[K, V]) (K, V) {
	panicIfNil[K, V](iter)
	if IsEmpty2(iter) {
		panic(emptyIterPanicMsg)
	}
	minK, minV := First2(iter)
	for k, v := range iter {
		if v < minV {
			minK, minV = k, v
		}
	}
	return minK, minV
}

// None returns true if none of the elements in the iter satisfy the predicate
// and false otherwise. It returns false at the first element that satisfies
// the predicate, and does not check the rest of the elements. If the iter is
// empty, it returns true. If the iter is nil, it panics.
func None[T any](iter iter.Seq[T], predicate Predicate[T]) bool {
	panicIfNil[any, T](iter)
	for v := range iter {
		if predicate(v) {
			return false
		}
	}
	return true
}

// None2 returns true if none of the key-value pairs in the iter satisfy the
// predicate and false otherwise. It returns false at the first element that
// satisfies the predicate, and does not check the rest of the elements. If
// the iter is empty, it returns true. If the iter is nil, it panics.
func None2[K comparable, V any](iter iter.Seq2[K, V], predicate Predicate2[K, V]) bool {
	panicIfNil[K, V](iter)
	for k, v := range iter {
		if predicate(k, v) {
			return false
		}
	}
	return true
}

// Product returns the product of all elements in the iter. If the iter is
// nil, it panics.
func Product[T types.Number](iter iter.Seq[T]) T {
	f := func(a, b T) T { return a * b }
	return Reduce(iter, f, T(1))
}

// Product2 returns the product of all values in the iter. If the iter is
// nil, it panics.
func Product2[K comparable, V types.Number](iter iter.Seq2[K, V]) V {
	f := func(a, b V, _ K) V { return a * b }
	return Reduce2(iter, f, V(1))
}

// Reduce applies the function to each element in the iter and returns the
// accumulated value. If the iter is nil, it panics.
func Reduce[T any](iter iter.Seq[T], f func(T, T) T, initialValue T) T {
	panicIfNil[any, T](iter)
	accumulated := initialValue
	for v := range iter {
		accumulated = f(accumulated, v)
	}
	return accumulated
}

// Reduce2 applies the function to each key-value pair in the iter and returns
// the accumulated value. If the iter is nil, it panics.
func Reduce2[K comparable, V any](iter iter.Seq2[K, V], f func(V, V, K) V, initialValue V) V {
	panicIfNil[K, V](iter)
	accumulated := initialValue
	for k, v := range iter {
		accumulated = f(accumulated, v, k)
	}
	return accumulated
}

// Sum returns the sum of all elements in the iter. If the iter is nil, it panics.
func Sum[T types.Summable](iter iter.Seq[T]) T {
	f := func(a, b T) T { return a + b }
	var zero T
	return Reduce(iter, f, zero)
}

// Sum2 returns the sum of all values in the iter. If the iter is nil, it panics.
func Sum2[K comparable, V types.Summable](iter iter.Seq2[K, V]) V {
	f := func(a, b V, _ K) V { return a + b }
	var zero V
	return Reduce2(iter, f, zero)
}
