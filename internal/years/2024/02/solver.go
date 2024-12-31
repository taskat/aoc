package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/intutils"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 2

// init registers the solver for day 02
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 02
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []report {
	reports := make([]report, len(lines))
	for i, line := range lines {
		reports[i] = parseReport(line)
	}
	return reports
}

// report is one line of report values
type report []int

// parseReport parses a report from a line
// The line is space separated integers
// In case of error, it panics
func parseReport(line string) report {
	parts := strings.Split(line, " ")
	values := slices.Map(parts, stringutils.Atoi)
	return report(values)
}

// isDecreasing checks if the report is strictly decreasing
func (r report) isDecreasing() bool {
	for i := 1; i < len(r); i++ {
		if r[i] >= r[i-1] {
			return false
		}
	}
	return true
}

// isDifferenceSafe checks if the report has no difference greater than 3
// between any two consecutive values
func (r report) isDifferenceSafe() bool {
	for i := 1; i < len(r); i++ {
		if intutils.Abs(r[i]-r[i-1]) > 3 {
			return false
		}
	}
	return true
}

// isIncreasing checks if the report is strictly increasing
func (r report) isIncreasing() bool {
	for i := 1; i < len(r); i++ {
		if r[i] <= r[i-1] {
			return false
		}
	}
	return true
}

// isSafe checks if the report is safe
// A report is safe if it is either decreasing or increasing
// and the difference between any two consecutive values is safe
func (r report) isSafe() bool {
	return (r.isDecreasing() || r.isIncreasing()) && r.isDifferenceSafe()
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	reports := s.parse(lines)
	safeCount := slices.Count(reports, report.isSafe)
	return strconv.Itoa(safeCount)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	reports := s.parse(lines)
	tolerableCount := slices.Count(reports, report.tolerable)
	return strconv.Itoa(tolerableCount)
}

// tolerable checks if the report is tolerable with the Problem Dampener.
// It means that removing any value from the report makes it safe
func (r report) tolerable() bool {
	for i := range r {
		dampened := report(slices.RemoveNth(r, i))
		if dampened.isSafe() {
			return true
		}
	}
	return false
}
