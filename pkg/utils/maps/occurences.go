package maps

// Occurrences creates a map with the occurrences of the elements in a slice.
func Occurrences[T comparable](s []T) map[T]int {
	m := make(map[T]int, len(s))
	for _, e := range s {
		m[e]++
	}
	return m
}

// AddOccurence adds an occurrence to a map.
func AddOccurence[T comparable](m map[T]int, e T, count int) {
	m[e] += count
}
