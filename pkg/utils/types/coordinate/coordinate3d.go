package coordinate

import (
	"fmt"

	m "math"

	"github.com/taskat/aoc/pkg/utils/math"
	"github.com/taskat/aoc/pkg/utils/types"
)

// Coordinate3D represents a 3D coordinate
type Coordinate3D[T types.Real] struct {
	X T
	Y T
	Z T
}

// NewCoordinate3D creates a new 3D coordinate
func NewCoordinate3D[T types.Real](x T, y T, z T) Coordinate3D[T] {
	return Coordinate3D[T]{X: x, Y: y, Z: z}
}

// Add adds the coordinates
func (c Coordinate3D[T]) Add(other Coordinate3D[T]) Coordinate3D[T] {
	return NewCoordinate3D(c.X+other.X, c.Y+other.Y, c.Z+other.Z)
}

// Distance calculates the Euclidean distance to another coordinate
func (c Coordinate3D[T]) Distance(other Coordinate3D[T]) float64 {
	return m.Sqrt(float64((c.X-other.X)*(c.X-other.X) +
		(c.Y-other.Y)*(c.Y-other.Y) +
		(c.Z-other.Z)*(c.Z-other.Z)))
}

// Equal checks if the coordinates are equal
func (c Coordinate3D[T]) Equal(other Coordinate3D[T]) bool {
	return c == other
}

// In3DSlice checks if the coordinate is in the slice
func (c Coordinate3D[T]) In3DSlice(width, height, depth T) bool {
	return c.InLimits(0, 0, 0, width-1, height-1, depth-1)
}

// InLimits checks if the coordinate is within the limits.
// All limits are inclusive.
func (c Coordinate3D[T]) InLimits(minX, minY, minZ, maxX, maxY, maxZ T) bool {
	return c.X >= minX && c.X <= maxX && c.Y >= minY && c.Y <= maxY && c.Z >= minZ && c.Z <= maxZ
}

// ManhattanDistance returns the Manhattan distance between the coordinates
func (c Coordinate3D[T]) ManhattanDistance(other Coordinate3D[T]) T {
	return math.Abs(c.X-other.X) + math.Abs(c.Y-other.Y) + math.Abs(c.Z-other.Z)
}

// String returns the string representation of the coordinate
func (c Coordinate3D[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", c.X, c.Y, c.Z)
}

// Sub subtracts the coordinates
func (c Coordinate3D[T]) Sub(other Coordinate3D[T]) Coordinate3D[T] {
	return NewCoordinate3D(c.X-other.X, c.Y-other.Y, c.Z-other.Z)
}
