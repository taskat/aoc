package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/intutils"
	"github.com/taskat/aoc/pkg/utils/slices"
)

// day is the day of the solver
const day = 01

// init registers the solver for day 01
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 01
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) ([]int, []int) {
	left, right := make([]int, len(lines)), make([]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "   ")
		left[i], _ = strconv.Atoi(parts[0])
		right[i], _ = strconv.Atoi(parts[1])
	}
	sort.Ints(left)
	sort.Ints(right)
	return left, right
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	left, right := s.parse(lines)
	distances := slices.Map_i(left, func(v int, i int) int {
		return intutils.Abs(v - right[i])
	})
	sum := slices.Sum(distances)
	return strconv.Itoa(sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	left, right := s.parse(lines)
	appearances := slices.Map(left, func(v int) int {
		return slices.Count(right, func(r int) bool {
			return r == v
		})
	})
	scores := slices.Map_i(appearances, func(v int, i int) int {
		return v * left[i]
	})
	sum := slices.Sum(scores)
	return strconv.Itoa(sum)
}
