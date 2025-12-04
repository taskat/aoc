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

// pattern represents a towel pattern design
type pattern string

// newPattern creates a new pattern from a string
func newPattern(s string) pattern {
	return pattern(s)
}

// String returns the string representation of the pattern
func (p pattern) String() string {
	return string(p)
}

// possibleCache contains the previously found possible patterns and subpatterns
var possibleCache = map[pattern]bool{"": true}

// clearCache clears the possible cache
func clearCache() {
	possibleCache = map[pattern]bool{"": true}
}

// isPossible checks if the pattern can be created from the available patterns
// It uses memoization to cache previously computed results
func (p pattern) isPossible(availablePatterns []pattern) bool {
	cacheResult, ok := possibleCache[p]
	if ok {
		return cacheResult
	}
	for _, available := range availablePatterns {
		remaining, found := strings.CutPrefix(p.String(), available.String())
		if !found {
			continue
		}
		remainingPossible := newPattern(remaining).isPossible(availablePatterns)
		possibleCache[p] = remainingPossible
		if remainingPossible {
			return true
		}
	}
	return false
}

// numberCache contains the previously computed number of possible ways to create patterns
var numberCache = map[pattern]int{"": 1}

// clearNumberCache clears the number cache
func clearNumberCache() {
	numberCache = map[pattern]int{"": 1}
}

// numberOfPossibleWays computes the number of possible ways to create the pattern
// from the available patterns. It uses memoization to cache previously computed results
func (p pattern) numberOfPossibleWays(availablePatterns []pattern) int {
	cacheResult, ok := numberCache[p]
	if ok {
		return cacheResult
	}
	count := 0
	for _, available := range availablePatterns {
		remaining, found := strings.CutPrefix(p.String(), available.String())
		if !found {
			continue
		}
		remainingPossbile := newPattern(remaining).numberOfPossibleWays(availablePatterns)
		count += remainingPossbile
	}
	numberCache[p] = count
	return count
}

// sortAvailable sorts the available patterns by length in ascending order
func sortAvailable(available []pattern) {
	slices2.SortFunc(available, func(a, b pattern) int {
		switch {
		case len(a) < len(b):
			return -1
		case len(a) > len(b):
			return 1
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
