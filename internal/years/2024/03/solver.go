package day03

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/slices"
)

// day is the day of the solver
const day = 3

// init registers the solver for day 03
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 03
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines, operations []string) []operation {
	r := strings.Join(operations, "|")
	mulRegexp := regexp.MustCompile(r)
	muls := make([]operation, 0)
	for _, line := range lines {
		matches := mulRegexp.FindAllString(line, -1)
		newMuls := slices.Map(matches, parseOperation)
		muls = append(muls, newMuls...)
	}
	return muls
}

type operation interface {
	eval(enable *bool) int
	regexp() string
}

// parseOperation parses the operation from a string
func parseOperation(s string) operation {
	switch {
	case strings.HasPrefix(s, "m"):
		return newMul(s)
	case strings.HasPrefix(s, "do("):
		return do{}
	case strings.HasPrefix(s, "don"):
		return dont{}
	}
	return nil
}

// do represents a do operation
type do struct{}

// eval evaluates the do operation
func (d do) eval(enable *bool) int {
	*enable = true
	return 0
}

// regexp returns the regular expression for the operation
func (do) regexp() string {
	return `do\(\)`
}

// dont represents a dont operation
type dont struct{}

// eval evaluates the dont operation
func (d dont) eval(enable *bool) int {
	*enable = false
	return 0
}

// regexp returns the regular expression for the operation
func (dont) regexp() string {
	return `don't\(\)`
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
func (m mul) eval(enable *bool) int {
	if !*enable {
		return 0
	}
	return m.left * m.right
}

// regexp returns the regular expression for the operation
func (mul) regexp() string {
	return `mul\(\d{1,3},\d{1,3}\)`
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	operations := s.parse(lines, []string{mul.regexp(mul{})})
	enable := true
	eval := func(op operation) int {
		return op.eval(&enable)
	}
	results := slices.Map(operations, eval)
	sum := slices.Sum(results)
	return fmt.Sprint(sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	operators := []string{mul.regexp(mul{}), do{}.regexp(), dont{}.regexp()}
	operations := s.parse(lines, operators)
	enable := true
	eval := func(op operation) int {
		return op.eval(&enable)
	}
	results := slices.Map(operations, eval)
	sum := slices.Sum(results)
	return fmt.Sprint(sum)
}
