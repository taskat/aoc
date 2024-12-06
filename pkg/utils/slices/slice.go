package slices

// Slice represents a generic slice type. It also includes the named types
type Slice[T any] interface {
	~[]T
}
