package iterutils

import (
	"cmp"
	"iter"
	"maps"
	"slices"
	"sort"
)

// Append appends the element to the input iterator. If the input iterator is nil, it panics.
func Append[T any](iter iter.Seq[T], v T) iter.Seq[T] {
	panicIfNil[any, T](iter)
	return func(yield func(T) bool) {
		for x := range iter {
			if !yield(x) {
				return
			}
		}
		yield(v)
	}
}

// Append2 appends the key-value pair to the input iterator. If the input iterator is nil, it panics.
func Append2[K comparable, V any](iter iter.Seq2[K, V], k K, v V) iter.Seq2[K, V] {
	panicIfNil[K, V](iter)
	return func(yield func(K, V) bool) {
		for key, value := range iter {
			if !yield(key, value) {
				return
			}
		}
		yield(k, v)
	}
}

// Concat concatenates the input iterators and returns a new iterator with the elements of the
// input iterators. If the input iterators are nil, it panics.
func Concat[T any](iters ...iter.Seq[T]) iter.Seq[T] {
	for _, iter := range iters {
		panicIfNil[any, T](iter)
	}
	return func(yield func(T) bool) {
		for _, iter := range iters {
			for v := range iter {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Concat2 concatenates the input iterators and returns a new iterator with the key-value pairs
// of the input iterators. If the input iterators are nil, it panics.
func Concat2[K comparable, V any](iters ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	for _, iter := range iters {
		panicIfNil[K, V](iter)
	}
	return func(yield func(K, V) bool) {
		for _, iter := range iters {
			for k, v := range iter {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Filter returns a new iterator with the elements that satisfy the predicate. If the
// input iterator is nil, it panics.
func Filter[T any](iter iter.Seq[T], predicate Predicate[T]) iter.Seq[T] {
	panicIfNil[any, T](iter)
	return func(yield func(T) bool) {
		for v := range iter {
			if predicate(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Filter2 returns a new iterator with the key-value pairs that satisfy the predicate.
// If the input iterator is nil, it panics.
func Filter2[K comparable, V any](iter iter.Seq2[K, V], predicate Predicate2[K, V]) iter.Seq2[K, V] {
	panicIfNil[K, V](iter)
	return func(yield func(K, V) bool) {
		for k, v := range iter {
			if predicate(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Keys returns a new iterator with the keys of the input iterator. If the input iterator
// is nil, it panics.
func Keys[K comparable, V any](iter iter.Seq2[K, V]) iter.Seq[K] {
	panicIfNil[K, V](iter)
	return func(yield func(K) bool) {
		for k := range iter {
			if !yield(k) {
				return
			}
		}
	}
}

// Map applies the function to each element of the input iterator and returns a new iterator
// with the results. If the input iterator is nil, it panics.
func Map[T, U any](iter iter.Seq[T], f func(T) U) iter.Seq[U] {
	panicIfNil[any, T](iter)
	return func(yield func(U) bool) {
		for v := range iter {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// Map12 applies the function to each element of the input iterator and returns a new iterator
// with the resulting key-value pairs. If the input iterator is nil, it panics.
func Map12[K comparable, V, U any](iter iter.Seq[V], f func(V) (K, U)) iter.Seq2[K, U] {
	panicIfNil[any, V](iter)
	return func(yield func(K, U) bool) {
		for v := range iter {
			k, u := f(v)
			if !yield(k, u) {
				return
			}
		}
	}
}

// Map2 applies the function to each key-value pair of the input iterator and returns a new
// iterator with the results. If the input iterator is nil, it panics.
func Map2[K, L comparable, V, U any](iter iter.Seq2[K, V], f func(K, V) (L, U)) iter.Seq2[L, U] {
	panicIfNil[K, V](iter)
	return func(yield func(L, U) bool) {
		for k, v := range iter {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

// Map21 applies the function to each key-value pair of the input iterator and returns a new
// iterator with the resulting values. If the input iterator is nil, it panics.
func Map21[K comparable, V, U any](iter iter.Seq2[K, V], f func(K, V) U) iter.Seq[U] {
	panicIfNil[K, V](iter)
	return func(yield func(U) bool) {
		for k, v := range iter {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

// RemoveKey returns a new iterator with the key-value pairs of the input iterator, removing
// the element with the given key. If the input iterator is nil, it panics.
func RemoveKey[K comparable, V any](iter iter.Seq2[K, V], key K) iter.Seq2[K, V] {
	panicIfNil[K, V](iter)
	return func(yield func(K, V) bool) {
		for k, v := range iter {
			if k != key {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// RemoveNth returns a new iterator with the elements of the input iterator, removing the
// element at the given index. If the input iterator is nil, it panics. If the index is out
// of bounds, the result will be the same as the input iterator.
func RemoveNth[T any](iter iter.Seq[T], index int) iter.Seq[T] {
	panicIfNil[any, T](iter)
	return func(yield func(T) bool) {
		i := 0
		for v := range iter {
			if i != index {
				if !yield(v) {
					return
				}
			}
			i++
		}
	}
}

// Skip returns a new iterator with the elements of the input iterator, skipping the first n
// elements. If the input iterator is nil, it panics. If n is greater than the length of the
// input iterator, the result will be an empty iterator.
func Skip[T any](iter iter.Seq[T], n int) iter.Seq[T] {
	panicIfNil[any, T](iter)
	return func(yield func(T) bool) {
		i := 0
		for v := range iter {
			if i >= n {
				if !yield(v) {
					return
				}
			}
			i++
		}
	}
}

// Skip2 returns a new iterator with the key-value pairs of the input iterator, skipping the
// first n elements. If the input iterator is nil, it panics. If n is greater than the length
// of the input iterator, the result will be an empty iterator.
func Skip2[K comparable, V any](iter iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	panicIfNil[K, V](iter)
	return func(yield func(K, V) bool) {
		i := 0
		for k, v := range iter {
			if i >= n {
				if !yield(k, v) {
					return
				}
			}
			i++
		}
	}
}

// Sort sorts the elements of the input iterator and returns a new iterator with the sorted
// elements. If the input iterator is nil, it panics.
func Sort[T cmp.Ordered](iter iter.Seq[T], less func(T, T) bool) iter.Seq[T] {
	panicIfNil[any, T](iter)
	return func(yield func(T) bool) {
		slice := slices.Collect(iter)
		sort.Slice(slice, func(i, j int) bool { return less(slice[i], slice[j]) })
		for _, v := range slice {
			if !yield(v) {
				return
			}
		}
	}
}

// Sort2 sorts the key-value pairs of the input iterator and returns a new iterator with the
// sorted key-value pairs. If the input iterator is nil, it panics.
func Sort2[K comparable, V cmp.Ordered](iter iter.Seq2[K, V], less func(K, V, K, V) bool) iter.Seq2[K, V] {
	panicIfNil[K, V](iter)
	return func(yield func(K, V) bool) {
		keys := slices.Collect(Keys(iter))
		values := maps.Collect(iter)
		sort.Slice(keys, func(i, j int) bool { return less(keys[i], values[keys[i]], keys[j], values[keys[j]]) })
		for _, k := range keys {
			if !yield(k, values[k]) {
				return
			}
		}
	}
}

// Split splits the elements of the input iterator into smaller iterators, based on the
// given predicate. If the input iterator is nil, it panics.
func Split[T any](i iter.Seq[T], predicate func(T) bool) iter.Seq[iter.Seq[T]] {
	panicIfNil[any, T](i)
	return func(yield func(iter.Seq[T]) bool) {
		var current []T
		for v := range i {
			if predicate(v) {
				if len(current) > 0 {
					if !yield(slices.Values(current)) {
						return
					}
				}
				current = nil
			} else {
				current = append(current, v)
			}
		}
		if len(current) > 0 {
			yield(slices.Values(current))
		}
	}
}

// Split2 splits the key-value pairs of the input iterator into smaller iterators, based on
// the given predicate. If the input iterator is nil, it panics.
func Split2[K comparable, V any](i iter.Seq2[K, V], predicate func(K, V) bool) iter.Seq[iter.Seq2[K, V]] {
	panicIfNil[K, V](i)
	return func(yield func(iter.Seq2[K, V]) bool) {
		var current map[K]V
		for k, v := range i {
			if predicate(k, v) {
				if current != nil {
					if !yield(maps.All(current)) {
						return
					}
				}
				current = nil
			} else {
				if current == nil {
					current = make(map[K]V)
				}
				current[k] = v
			}
		}
		if current != nil {
			yield(maps.All(current))
		}
	}
}

// Swap swaps the elements at the given indices in the input iterator. If the input iterator
// is nil, it panics. If the indices are out of bounds, it panics.
func Swap[T any](iter iter.Seq[T], i, j int) iter.Seq[T] {
	panicIfNil[any, T](iter)
	if i < 0 || j < 0 || i >= Len(iter) || j >= Len(iter) {
		panic(indexOOBPanicMsg)
	}
	return func(yield func(T) bool) {
		slice := slices.Collect(iter)
		slice[i], slice[j] = slice[j], slice[i]
		for _, v := range slice {
			if !yield(v) {
				return
			}
		}
	}
}

// Swap2 swaps the key-value pairs at the given indices in the input iterator. If the input
// iterator is nil, it panics. If the indices are out of bounds, it panics.
func Swap2[K comparable, V any](iter iter.Seq2[K, V], i, j int) iter.Seq2[K, V] {
	panicIfNil[K, V](iter)
	if i < 0 || j < 0 || i >= Len2(iter) || j >= Len2(iter) {
		panic(indexOOBPanicMsg)
	}
	return func(yield func(K, V) bool) {
		keys := slices.Collect(Keys(iter))
		values := maps.Collect(iter)
		keys[i], keys[j] = keys[j], keys[i]
		values[keys[i]], values[keys[j]] = values[keys[j]], values[keys[i]]
		for _, k := range keys {
			if !yield(k, values[k]) {
				return
			}
		}
	}
}

// SwapByKey swaps the values of the key-value pairs with the given keys in the input iterator.
// If the input iterator is nil, it panics. If the keys are not present in the input iterator,
// it panics.
func SwapByKey[K comparable, V any](iter iter.Seq2[K, V], key1, key2 K) iter.Seq2[K, V] {
	panicIfNil[K, V](iter)
	keys := Keys(iter)
	if !Contains(keys, key1) || !Contains(keys, key2) {
		panic(keyNotFoundPanicMsg)
	}
	keysIter := slices.All(slices.Collect(keys))
	index1, _, _ := Find2(keysIter, func(_ int, v K) bool { return v == key1 })
	index2, _, _ := Find2(keysIter, func(_ int, v K) bool { return v == key2 })
	keys = Swap(keys, index1, index2)
	all := maps.Collect(iter)
	return func(yield func(K, V) bool) {
		for k := range keys {
			if !yield(k, all[k]) {
				return
			}
		}
	}
}

// Values returns a new iterator with the values of the input iterator. If the input iterator
// is nil, it panics.
func Values[K comparable, V any](iter iter.Seq2[K, V]) iter.Seq[V] {
	panicIfNil[K, V](iter)
	return func(yield func(V) bool) {
		for _, v := range iter {
			if !yield(v) {
				return
			}
		}
	}
}

// Zip applies the function to each pair of elements from the two input iterators and returns
// a new iterator with the results. If the input iterators are nil, it panics. If the input iterators
// have different lengths, the result will have the length of the shortest iterator.
func Zip[T, U, V any](iter1 iter.Seq[T], iter2 iter.Seq[U], f func(T, U) V) iter.Seq[V] {
	panicIfNil[any, T](iter1)
	panicIfNil[any, U](iter2)
	return func(yield func(V) bool) {
		next1, stop1 := iter.Pull(iter1)
		defer stop1()
		next2, stop2 := iter.Pull(iter2)
		defer stop2()
		for {
			v1, ok1 := next1()
			v2, ok2 := next2()
			if !ok1 || !ok2 {
				return
			}
			if !yield(f(v1, v2)) {
				return
			}
		}
	}
}
