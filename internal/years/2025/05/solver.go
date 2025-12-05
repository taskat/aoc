package day05

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
	rangetype "github.com/taskat/aoc/pkg/utils/types/range"
)

// day is the day of the solver
const day = 5

// init registers the solver for day 05
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 05
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) ([]rangetype.Range, []int) {
	delimIdx := slices.FindIndex(lines, stringutils.IsEmpty)
	ranges := slices.Map(lines[:delimIdx], parseRange)
	ids := slices.Map(lines[delimIdx+1:], stringutils.Atoi)
	return ranges, ids
}

// parseRange parses a single line into an all inclusive range
func parseRange(line string) rangetype.Range {
	var start, end int
	fmt.Sscanf(strings.TrimSpace(line), "%d-%d", &start, &end)
	return rangetype.NewAllInclusive(start, end)
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	ranges, ids := s.parse(lines)
	countFresh := 0
	for _, id := range ids {
		for i, r := range ranges {
			if r.Contains(id) {
				countFresh++
				break
			}
			_ = i
		}
	}
	return fmt.Sprintf("%d", countFresh)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
