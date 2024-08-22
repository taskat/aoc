package day1

import (
	"fmt"
	"taskat/aoc/internal/years/2023/days"
)

// day is the day of the solver
const day = 1

// init registers the solver for day 1
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 1
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// ParsePart1 parses the input for part 1
func (s *Solver) ParsePart1(lines []string) {
	s.parse(lines)
}

// ParsePart2 parses the input for part 2
func (s *Solver) ParsePart2(lines []string) {
	s.parse(lines)
}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) {}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1() string {
	return "a"
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2() string {
	return "b"
}
