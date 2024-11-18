package year2023

import (
	"github.com/taskat/aoc/cmd/main/years"
	"github.com/taskat/aoc/internal/years/2023/days"
	_ "github.com/taskat/aoc/internal/years/2023/imports"
)

// init register the year's collection of solvers
func init() {
	years.AddYear(days.Year, days.Current)
}
