package day19

import (
	"fmt"
	slices2 "slices"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/slices"
)

// day is the day of the solver
const day = 19

// init registers the solver for day 19
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 19
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) ([]pattern, []pattern) {
	parts := strings.Split(lines[0], ", ")
	available := slices.Map(parts, newPattern)
	desired := slices.Map(lines[2:], newPattern)
	return available, desired
}

type pattern string

func newPattern(s string) pattern {
	return pattern(s)
}

var possibleCache = map[pattern]bool{"": true}

func clearCache() {
	possibleCache = map[pattern]bool{"": true}
}

func (p pattern) isPossible(available []pattern) bool {
	cacheResult, ok := possibleCache[p]
	if ok {
		return cacheResult
	}
	for _, a := range available {
		if strings.HasPrefix(string(p), string(a)) {
			remainingPossbile := newPattern(strings.TrimPrefix(string(p), string(a))).isPossible(available)
			possibleCache[p] = remainingPossbile
			if remainingPossbile {
				return true
			}
		}
	}
	return false
}

var numberCache = map[pattern]int{"": 1}

func clearNumberCache() {
	numberCache = map[pattern]int{"": 1}
}

func (p pattern) numberOfPossibleWays(available []pattern) int {
	cacheResult, ok := numberCache[p]
	if ok {
		return cacheResult
	}
	count := 0
	for _, a := range available {
		if strings.HasPrefix(string(p), string(a)) {
			remainingPossbile := newPattern(strings.TrimPrefix(string(p), string(a))).numberOfPossibleWays(available)
			count += remainingPossbile
		}
	}
	numberCache[p] = count
	return count
}

func sortAvailable(available []pattern) {
	slices2.SortFunc(available, func(a, b pattern) int {
		switch {
		case len(a) < len(b):
			return 1
		case len(a) > len(b):
			return -1
		default:
			return 0
		}
	})
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	available, desired := s.parse(lines)
	sortAvailable(available)
	clearCache()
	possible := slices.Filter(desired, func(p pattern) bool {
		return p.isPossible(available)
	})
	return fmt.Sprint(len(possible))
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	available, desired := s.parse(lines)
	sortAvailable(available)
	clearNumberCache()
	counts := slices.Map(desired, func(p pattern) int {
		return p.numberOfPossibleWays(available)
	})
	return fmt.Sprint(slices.Sum(counts))
}
