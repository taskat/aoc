package days

import (
	"taskat/aoc/cmd/main/collection"
	"taskat/aoc/cmd/main/solver"
)

// Current is the collection of solvers for the year 2023
var Current = make(collection.Year)

// AddDay adds a new solver to the current collection
func AddDay(day int, s solver.Solver) {
	Current[day] = s
}
