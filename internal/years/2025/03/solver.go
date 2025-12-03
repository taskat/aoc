package day03

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 3

// init registers the solver for day 03
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 03
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []bank {
	return slices.Map(lines, parseBank)
}

// bank represents a bank of batteries
type bank []int

// maxJoltage returns the maximum joltage in the bank
// It is the largest value, that is concatenated with n batteries in order
func (b bank) maxJoltage(n int) int {
	digits := make([]int, n)
	from := 0
	to := len(b) - n + 1
	for i := range n {
		var digitIdx int
		digits[i], digitIdx = slices.Max_i(b[from:to])
		from += digitIdx + 1
		to++
	}
	return slices.Reduce(digits, func(acc, next int) int { return 10*acc + next }, 0)
}

// parseBank parses a bank of batteries from a string
func parseBank(line string) bank {
	b := make(bank, len(line))
	for i, ch := range line {
		b[i] = stringutils.RuneToInt(ch)
	}
	return b
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	banks := s.parse(lines)
	joltages := slices.Map(banks, func(b bank) int { return b.maxJoltage(2) })
	sum := slices.Sum(joltages)
	return fmt.Sprint(sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	banks := s.parse(lines)
	joltages := slices.Map(banks, func(b bank) int { return b.maxJoltage(12) })
	sum := slices.Sum(joltages)
	return fmt.Sprint(sum)
}
