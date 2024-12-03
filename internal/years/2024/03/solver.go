package day03

import (
	"fmt"
	"regexp"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/slices"
)

// day is the day of the solver
const day = 03

// init registers the solver for day 03
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 03
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []mul {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	muls := make([]mul, 0)
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		newMuls := slices.Map(matches, newMul)
		muls = append(muls, newMuls...)
	}
	return muls
}

// mul represents a multiplication operation
type mul struct {
	left  int
	right int
}

// newMul creates a new multiplication operation
func newMul(s string) mul {
	left, right := 0, 0
	fmt.Sscanf(s, "mul(%d,%d)", &left, &right)
	return mul{left: left, right: right}
}

// eval evaluates the multiplication operation
func (m mul) eval() int {
	return m.left * m.right
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	muls := s.parse(lines)
	results := slices.Map(muls, mul.eval)
	sum := slices.Sum(results)
	return fmt.Sprint(sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
