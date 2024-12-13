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
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 13
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []machine {
	var machines []machine
	for i := 0; i < len(lines); i += 4 {
		machines = append(machines, parseMachine(lines[i:i+3]))
	}
	return machines
}

type machine struct {
	buttonA coord
	buttonB coord
	prize   coord
}

func parseMachine(lines []string) machine {
	return machine{
		buttonA: parseCoord(lines[0], "Button A: X+%d, Y+%d"),
		buttonB: parseCoord(lines[1], "Button B: X+%d, Y+%d"),
		prize:   parseCoord(lines[2], "Prize: X=%d, Y=%d"),
	}
}

func (m *machine) addOffset(offset int) {
	m.prize.X += offset
	m.prize.Y += offset
}

func (m machine) cost(a, b int) int {
	return 3*a + b
}

func (m machine) tokens() int {
	a, b := m.solve()
	if a == 0 && b == 0 {
		return 0
	}
	return m.cost(a, b)
}

func (m machine) solve() (int, int) {
	nomB := m.buttonA.X*m.prize.Y - m.buttonA.Y*m.prize.X
	denomB := m.buttonA.X*m.buttonB.Y - m.buttonA.Y*m.buttonB.X
	if nomB%denomB != 0 {
		return 0, 0
	}
	b := nomB / denomB
	if b < 0 {
		return 0, 0
	}
	nomA := m.prize.X - m.buttonB.X*b
	denomA := m.buttonA.X
	if nomA%denomA != 0 {
		return 0, 0
	}
	a := nomA / denomA
	if a < 0 {
		return 0, 0
	}
	return a, b
}

type coord = coordinate.Coordinate2D[int]

func parseCoord(line string, format string) coord {
	var x, y int
	fmt.Sscanf(line, format, &x, &y)
	return coordinate.NewCoordinate2D(x, y)
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	machines := s.parse(lines)
	tokens := slices.Map(machines, machine.tokens)
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
	tokens := slices.Map(machines, machine.tokens)
	sum := slices.Sum(tokens)
	return fmt.Sprint(sum)
}
