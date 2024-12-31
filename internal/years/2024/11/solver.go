package day11

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/intutils"
	"github.com/taskat/aoc/pkg/utils/maps"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 11

// init registers the solver for day 11
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 11
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) stones {
	numbers := slices.Map(strings.Split(lines[0], " "), stringutils.Atoi)
	return maps.Occurrences(numbers)
}

// stones contains the value of the stones and the count of each
// value
type stones map[int]int

// blink evaluates the next state of the stones and updates the
// current state
func (s *stones) blink() {
	newStones := make(stones)
	maps.ForEach(*s, func(stone, count int) {
		values := nextValues(stone)
		slices.ForEach(values, func(value int) {
			maps.AddOccurence(newStones, value, count)
		})
	})
	*s = newStones
}

// count returns the total count of stones
func (s stones) count() int {
	return maps.Sum(s)
}

// simulate runs the blink for n times
func (s *stones) simulate(n int) {
	for i := 0; i < n; i++ {
		s.blink()
	}
}

// nextValues returns the next values based on the current value of the stone
func nextValues(value int) [](int) {
	valueLength := intutils.Length(value)
	switch {
	case value == 0:
		return []int{1}
	case valueLength%2 == 0:
		divider := intutils.Power(10, valueLength/2)
		return []int{value / divider, value % divider}
	default:
		return []int{value * 2024}
	}
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	stones := s.parse(lines)
	stones.simulate(25)
	return fmt.Sprint(stones.count())
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	stones := s.parse(lines)
	stones.simulate(75)
	return fmt.Sprint(stones.count())
}
