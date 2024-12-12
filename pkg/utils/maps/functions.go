package maps

import "github.com/taskat/aoc/pkg/utils/types"

// Append appends the element to the slice in the map. If the key is not present,
// a new slice is created. If the key is present, the element is appended to the
// existing slice. If the map is nil, it panics.
func Append[K comparable, V any](m map[K][]V, k K, v V) {
	if m == nil {
		panic("map is nil")
	}
	if !Contains(m, k) {
		m[k] = make([]V, 0)
	}
	m[k] = append(m[k], v)
}

// Contains checks if a map contains a key.
func Contains[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

// ForEach iterates over the elements of a map.
func ForEach[K comparable, V any](m map[K]V, f func(K, V)) {
	for k, v := range m {
		f(k, v)
	}
}

// Keys returns the keys of a map.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Map maps the values of a map.
func Map[K1, K2 comparable, V1, V2 any](m map[K1]V1, f func(K1, V1) (K2, V2)) map[K2]V2 {
	result := make(map[K2]V2, len(m))
	for k, v := range m {
		k2, v2 := f(k, v)
		result[k2] = v2
	}
	return result
}

// MapValues maps the values of a map.
func MapValues[K comparable, V1, V2 any](m map[K]V1, f func(V1) V2) map[K]V2 {
	result := make(map[K]V2, len(m))
	for k, v := range m {
		result[k] = f(v)
	}
	return result
}

// Merge merges two maps into a new map. If a key is present in both maps,
// the value from the second map is used.
func Merge[K comparable, V any](m1, m2 map[K]V) map[K]V {
	result := make(map[K]V, len(m1)+len(m2))
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	return result
}

// Sum sums the values of a map.
func Sum[K comparable, V types.Summable](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

// Values returns the values of a map.
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
