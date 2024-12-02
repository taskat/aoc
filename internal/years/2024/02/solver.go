package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/intutils"
	"github.com/taskat/aoc/pkg/utils/slices"
)

// day is the day of the solver
const day = 02

// init registers the solver for day 02
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 02
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []report {
	reports := make([]report, len(lines))
	for i, line := range lines {
		reports[i] = parseReport(line)
	}
	return reports
}

type report []int

func parseReport(line string) report {
	parts := strings.Split(line, " ")
	report := make(report, len(parts))
	var err error
	for i, part := range parts {
		report[i], err = strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
	}
	return report
}

func (r report) isDecreasing() bool {
	for i := 1; i < len(r); i++ {
		if r[i] >= r[i-1] {
			return false
		}
	}
	return true
}

func (r report) isDifferenceSafe() bool {
	for i := 1; i < len(r); i++ {
		if intutils.Abs(r[i]-r[i-1]) > 3 {
			return false
		}
	}
	return true
}

func (r report) isIncreasing() bool {
	for i := 1; i < len(r); i++ {
		if r[i] <= r[i-1] {
			return false
		}
	}
	return true
}

func (r report) isSafe() bool {
	return (r.isDecreasing() || r.isIncreasing()) && r.isDifferenceSafe()
}

func (r report) tolerable() bool {
	for i := range r {
		dampened := make(report, len(r)-1)
		copy(dampened[:i], r[:i])
		copy(dampened[i:], r[i+1:])
		if dampened.isSafe() {
			return true
		}
	}
	return false
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
