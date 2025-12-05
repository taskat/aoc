package rangetype

import "fmt"

// Range represents a generic range type. It contains a start and end value.
// The start is inclusive, and the end is exclusive.
type Range struct {
	Start int
	End   int
}

// New creates a new Range instance. If end is less than start, it panics.
func New(start, end int) Range {
	if end <= start {
		panic("end must be greater than to start")
	}
	return Range{Start: start, End: end}
}

// NewAllInclusive creates a new Range instance that includes both start and end.
func NewAllInclusive(start, end int) Range {
	return New(start, end+1)
}

// NewAllExclusive creates a new Range instance that excludes both start and end.
func NewAllExclusive(start, end int) Range {
	return New(start+1, end)
}

// Contains checks if a value is within the range.
func (r Range) Contains(value int) bool {
	return value >= r.Start && value < r.End
}

// Length returns the length of the range.
func (r Range) Length() int {
	return r.End - r.Start
}

// Overlaps checks if two ranges overlap.
func (r Range) Overlaps(other Range) bool {
	return r.Start < other.End && other.Start < r.End
}

// String returns a string representation of the range.
func (r Range) String() string {
	return fmt.Sprintf("[%d, %d)", r.Start, r.End)
}
