package day11

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/intutils"
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
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) stones {
	stones := make(stones)
	numbers := slices.Map(strings.Split(lines[0], " "), stringutils.Atoi)
	for _, n := range numbers {
		stones[n]++
	}
	return stones
}

type stones map[int]int

func (s *stones) blink() {
	newStones := make(stones)
	for stone, count := range *s {
		values := nextValues(stone)
		for _, v := range values {
			newStones[v] += count
		}
	}
	*s = newStones
}

func (s stones) count() int {
	count := 0
	for _, c := range s {
		count += c
	}
	return count
}

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
	for i := 0; i < 25; i++ {
		stones.blink()
	}
	return fmt.Sprint(stones.count())
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
