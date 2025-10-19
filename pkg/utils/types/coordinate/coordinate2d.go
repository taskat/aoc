package coordinate

import (
	"fmt"

	"github.com/taskat/aoc/pkg/utils/math"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/types"
)

// Coordinate2D represents a 2D coordinate
type Coordinate2D[T types.Real] struct {
	X T
	Y T
}

// FromIndexes creates a new 2D coordinate from the indexes
func FromIndexes[T types.Integer](i, j T) Coordinate2D[T] {
	return NewCoordinate2D(j, -i)
}

// NewCoordinate2D creates a new 2D coordinate
func NewCoordinate2D[T types.Real](x T, y T) Coordinate2D[T] {
	return Coordinate2D[T]{X: x, Y: y}
}

// Add adds the coordinates
func (c Coordinate2D[T]) Add(other Coordinate2D[T]) Coordinate2D[T] {
	return NewCoordinate2D(c.X+other.X, c.Y+other.Y)
}

// Column returns the column of the coordinate.
// It panics if not used with 2D slices.
func (c Coordinate2D[T]) Column() T {
	if !c.Is2DSliceCoord() {
		panic(fmt.Sprintf("%s is not a 2D slice coordinate", c))
	}
	return c.X
}

// Equal checks if the coordinates are equal
func (c Coordinate2D[T]) Equal(other Coordinate2D[T]) bool {
	return c == other
}

// Go returns the coordinate after moving in the direction
func (c Coordinate2D[T]) Go(direction Direction) Coordinate2D[T] {
	return c.GoN(direction, 1)
}

// GoN returns the coordinate after moving n times in the direction
func (c Coordinate2D[T]) GoN(direction Direction, n int) Coordinate2D[T] {
	vec := direction.ToCoordinate2D()
	return c.Add(NewCoordinate2D(T(vec.X*n), T(vec.Y*n)))
}

// I returns the Y coordinate. It is used for 2D slices
// (indexes are row-major, so I is Y). It panics if not used with 2D slices.
func (c Coordinate2D[T]) I() T {
	if !c.Is2DSliceCoord() {
		panic(fmt.Sprintf("%s is not a 2D slice coordinate", c))
	}
	return -c.Y
}

// In2DSlice checks if the coordinate is in the slice
func (c Coordinate2D[T]) In2DSlice(width, height T) bool {
	return c.InLimits(0, -height+1, width-1, 0)
}

// InLimits checks if the coordinate is within the limits.
// All limits are inclusive.
func (c Coordinate2D[T]) InLimits(minX, minY, maxX, maxY T) bool {
	return c.X >= minX && c.X <= maxX && c.Y >= minY && c.Y <= maxY
}

// Is2DSliceCoord checks if the coordinate is a valid index in a 2D slice
// (indexes are row-major, so Y is negative)
func (c Coordinate2D[T]) Is2DSliceCoord() bool {
	return c.Y <= 0 && c.X >= 0
}

// J returns the X coordinate. It is used for 2D slices
// (indexes are row-major, so J is X). It panics if not used with 2D slices.
func (c Coordinate2D[T]) J() T {
	if !c.Is2DSliceCoord() {
		panic(fmt.Sprintf("%s is not a 2D slice coordinate", c))
	}
	return c.X
}

// ManhattanDistance returns the Manhattan distance between the coordinates
func (c Coordinate2D[T]) ManhattanDistance(other Coordinate2D[T]) T {
	return math.Abs(c.X-other.X) + math.Abs(c.Y-other.Y)
}

// Move moves the coordinate in the given direction
// It modifies the coordinate in place
func (c *Coordinate2D[T]) Move(direction Direction) {
	c.MoveN(direction, 1)
}

// MoveN moves the coordinate n times in the given direction
// It modifies the coordinate in place
func (c *Coordinate2D[T]) MoveN(direction Direction, n int) {
	vec := direction.ToCoordinate2D()
	c.X += T(vec.X * n)
	c.Y += T(vec.Y * n)
}

// Neighbors returns the neighbors of the coordinate in the given directions
func (c Coordinate2D[T]) Neighbors(directions []Direction) map[Direction]Coordinate2D[T] {
	coords := slices.Map(directions, c.Go)
	return slices.ToMap(directions, coords)
}

// Row returns the row of the coordinate
// It panics if not used with 2D slices.
func (c Coordinate2D[T]) Row() T {
	if !c.Is2DSliceCoord() {
		panic(fmt.Sprintf("%s is not a 2D slice coordinate", c))
	}
	return -c.Y
}

// String returns the string representation of the coordinate
func (c Coordinate2D[T]) String() string {
	return fmt.Sprintf("(%v, %v)", c.X, c.Y)
}

// Sub subtracts the coordinates
func (c Coordinate2D[T]) Sub(other Coordinate2D[T]) Coordinate2D[T] {
	return NewCoordinate2D(c.X-other.X, c.Y-other.Y)
}
