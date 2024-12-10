package coordinate

import "github.com/taskat/aoc/pkg/utils/types"

// Coordinate2D represents a 2D coordinate
type Coordinate2D[T types.Real] struct {
	X T
	Y T
}

// NewCoordinate2D creates a new 2D coordinate
func NewCoordinate2D[T types.Real](x T, y T) Coordinate2D[T] {
	return Coordinate2D[T]{X: x, Y: y}
}

// Add adds the coordinates
func (c Coordinate2D[T]) Add(other Coordinate2D[T]) Coordinate2D[T] {
	return NewCoordinate2D(c.X+other.X, c.Y+other.Y)
}

// Go returns the coordinate after moving in the direction
func (c Coordinate2D[T]) Go(direction Direction) Coordinate2D[T] {
	vec := direction.ToCoordinate2D()
	return c.Add(NewCoordinate2D(T(vec.X), T(vec.Y)))
}

// In2DSlice checks if the coordinate is in the slice
func (c Coordinate2D[T]) In2DSlice(width, height T) bool {
	return c.InLimits(0, 0, width-1, height-1)
}

// InLimits checks if the coordinate is within the limits.
// All limits are inclusive.
func (c Coordinate2D[T]) InLimits(minX, minY, maxX, maxY T) bool {
	return c.X >= minX && c.X <= maxX && c.Y >= minY && c.Y <= maxY
}
