package year2023

import (
	"taskat/aoc/cmd/main/years"
	"taskat/aoc/internal/years/2023/days"
	_ "taskat/aoc/internal/years/2023/imports"
)

// year is the year of the solvers
const year = 2023

// init register the year's collection of solvers
func init() {
	years.AddYear(year, days.Current)
}
