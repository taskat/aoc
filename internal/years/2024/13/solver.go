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

func (m machine) cost(a, b int) int {
	return 3*a + b
}

func (m machine) fewestTokens() result {
	min := 400
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			if m.possible(a, b) {
				cost := m.cost(a, b)
				if cost < min {
					min = cost
				}
			}
		}
	}
	if min == 400 {
		return newResult(0, false)
	}
	return newResult(min, true)
}

func (m machine) possible(a, b int) bool {
	xMacthes := m.buttonA.X*a+m.buttonB.X*b == m.prize.X
	yMatches := m.buttonA.Y*a+m.buttonB.Y*b == m.prize.Y
	return xMacthes && yMatches
}

type coord = coordinate.Coordinate2D[int]

func parseCoord(line string, format string) coord {
	var x, y int
	fmt.Sscanf(line, format, &x, &y)
	return coordinate.NewCoordinate2D(x, y)
}

type result struct {
	tokens   int
	possible bool
}

func newResult(tokens int, possible bool) result {
	return result{tokens, possible}
}

func (r result) isPossible() bool {
	return r.possible
}

func (r result) minTokens() int {
	return r.tokens
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	machines := s.parse(lines)
	results := slices.Map(machines, machine.fewestTokens)
	results = slices.Filter(results, result.isPossible)
	tokens := slices.Map(results, result.minTokens)
	sum := slices.Sum(tokens)
	return fmt.Sprint(sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
