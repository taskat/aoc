package coordinate

import "github.com/taskat/aoc/pkg/utils/types"

// Actor represents an entity that can act on a 2D coordinate grid.
// It has a position and can move in specified directions.
type Actor[T types.Real] struct {
	Position  Coordinate2D[T]
	Direction Direction
}

// NewActor creates a new Actor at the given position and direction.
func NewActor[T types.Real](position Coordinate2D[T], direction Direction) *Actor[T] {
	return &Actor[T]{
		Position:  position,
		Direction: direction,
	}
}

// Move moves the actor in the direction it's facing.
func (a *Actor[T]) Move() {
	a.Position = a.Position.Go(a.Direction)
}

// Next returns the next position of the actor based on its current direction.
func (a Actor[T]) Next() Coordinate2D[T] {
	return a.Position.Go(a.Direction)
}
