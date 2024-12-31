package day13

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 13

// init registers the solver for day 13
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 13
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []machine {
	var machines []machine
	for i := 0; i < len(lines); i += 4 {
		machines = append(machines, parseMachine(lines[i:i+3]))
	}
	return machines
}

// machine a claw machine configuration
type machine struct {
	buttonA coord
	buttonB coord
	prize   coord
}

// parseMachine parses a machine from the input lines
// The input lines are expected to be in the following format:
// Button A: X+%d, Y+%d
// Button B: X+%d, Y+%d
// Prize: X=%d, Y=%d
func parseMachine(lines []string) machine {
	return machine{
		buttonA: parseCoord(lines[0], "Button A: X+%d, Y+%d"),
		buttonB: parseCoord(lines[1], "Button B: X+%d, Y+%d"),
		prize:   parseCoord(lines[2], "Prize: X=%d, Y=%d"),
	}
}

// addOffset adds an offset to the prize coordinates
func (m *machine) addOffset(offset int) {
	m.prize.X += offset
	m.prize.Y += offset
}

// cost calculates the cost of the machine
func (m machine) cost() int {
	a, b := m.solve()
	if a == 0 && b == 0 {
		return 0
	}
	return a*3 + b
}

// solve solves the equation system of the machine. If there is a solution,
// which is a positive integer for both variables, it returns the values.
// Otherwise, it returns 0, 0.
func (m machine) solve() (int, int) {
	nomB := m.buttonA.X*m.prize.Y - m.buttonA.Y*m.prize.X
	denomB := m.buttonA.X*m.buttonB.Y - m.buttonA.Y*m.buttonB.X
	if !isPositiveInteger(nomB, denomB) {
		return 0, 0
	}
	b := nomB / denomB
	nomA := m.prize.X - m.buttonB.X*b
	denomA := m.buttonA.X
	if !isPositiveInteger(nomA, denomA) {
		return 0, 0
	}
	a := nomA / denomA
	return a, b
}

// isPositiveInteger checks if the given fraction is a positive integer
func isPositiveInteger(nom, denom int) bool {
	return nom%denom == 0 && nom/denom >= 0
}

// coord is  shorthand for a 2D coordinate
type coord = coordinate.Coordinate2D[int]

// parseCoord parses a coordinate from a line
func parseCoord(line string, format string) coord {
	var x, y int
	fmt.Sscanf(line, format, &x, &y)
	return coordinate.NewCoordinate2D(x, y)
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	machines := s.parse(lines)
	tokens := slices.Map(machines, machine.cost)
	sum := slices.Sum(tokens)
	return fmt.Sprint(sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	machines := s.parse(lines)
	offset := 10_000_000_000_000
	addOffset := func(m machine) machine {
		m.addOffset(offset)
		return m
	}
	machines = slices.Map(machines, addOffset)
	tokens := slices.Map(machines, machine.cost)
	sum := slices.Sum(tokens)
	return fmt.Sprint(sum)
}
