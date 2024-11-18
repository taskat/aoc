package collection

import "github.com/taskat/aoc/cmd/main/solver"

// YearCollection is a map that contains all the solvers for a specific year
type Collection[T any] map[int]T

// Add adds a new element to the collection
func (c Collection[T]) Add(key int, elem T) {
	c[key] = elem
}

// Get returns the element corresponding to the key
func (c Collection[T]) Get(key int) T {
	return c[key]
}

// Year is a collection of days' solvers
type Year = Collection[solver.Solver]
