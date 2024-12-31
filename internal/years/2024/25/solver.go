package day25

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 25

// init registers the solver for day 25
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 25
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

type schematic []int

func (s schematic) overlap(other schematic) bool {
	for i, h := range s {
		if h+other[i] > 5 {
			return true
		}
	}
	return false
}

func parseKey(lines []string) schematic {
	heights := make(schematic, len(lines[0]))
	for j := 0; j < len(lines[0]); j++ {
		for i := len(lines) - 1; i >= 0; i-- {
			if lines[i][j] == '.' {
				heights[j] = len(lines) - i - 2
				break
			}
		}
	}
	return heights
}

func parseLock(lines []string) schematic {
	heights := make(schematic, len(lines[0]))
	for j := 0; j < len(lines[0]); j++ {
		for i := 0; i < len(lines); i++ {
			if lines[i][j] == '.' {
				heights[j] = i - 1
				break
			}
		}
	}
	return heights
}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) ([]schematic, []schematic) {
	keys := make([]schematic, 0)
	locks := make([]schematic, 0)
	blocks := slices.Split(lines, stringutils.IsEmpty)
	for _, block := range blocks {
		if block[0][0] == '#' {
			locks = append(locks, parseLock(block))
		} else {
			keys = append(keys, parseKey(block))
		}
	}
	return keys, locks
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	keys, locks := s.parse(lines)
	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			if !lock.overlap(key) {
				count++
			}
		}
	}
	return fmt.Sprint(count)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
