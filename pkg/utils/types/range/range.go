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

// Merge merges the other range into the current range if they overlap or are adjacent.
// It returns true if the merge was successful, false otherwise.
func (r *Range) Merge(other Range) bool {
	if !r.Overlaps(other) && r.End != other.Start && other.End != r.Start {
		return false
	}
	if other.Start < r.Start {
		r.Start = other.Start
	}
	if other.End > r.End {
		r.End = other.End
	}
	return true
}

// Overlaps checks if two ranges overlap.
func (r Range) Overlaps(other Range) bool {
	return r.Start < other.End && other.Start < r.End
}

// String returns a string representation of the range.
func (r Range) String() string {
	return fmt.Sprintf("[%d, %d)", r.Start, r.End)
}
