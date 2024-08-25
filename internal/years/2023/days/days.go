package days

import (
	"taskat/aoc/cmd/main/collection"
	"taskat/aoc/cmd/main/solver"
)

// Year is the year of the solvers
const Year = 2023

// Current is the collection of solvers for the year 2023
var Current = make(collection.Year)

// AddDay adds a new solver to the current collection
func AddDay(day int, s solver.Solver) {
	Current[day] = s
}
