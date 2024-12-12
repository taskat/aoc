package set

// Set is a set. It is implemented as a map with empty structs.
type Set[T comparable] map[T]struct{}

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

// ToSlice converts the set to a slice.
func (s Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s))
	for k := range s {
		result = append(result, k)
	}
	return result
}
