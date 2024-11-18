package day01

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2023/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 01

// init registers the solver for day 01
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 01
type Solver struct {
	lines []string
}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// ParsePart1 parses the input for part 1
func (s *Solver) ParsePart1(lines []string) {
	s.parse(lines)
}

// ParsePart2 parses the input for part 2
func (s *Solver) ParsePart2(lines []string) {
	s.parse(lines)
}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) {
	s.lines = lines
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1() string {
	calibrationValues := slices.Map(s.lines, getCalibrationValue)
	sum := slices.Sum(calibrationValues)
	return fmt.Sprintf("%d", sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2() string {
	return ""
}

func getDigits(line string) []int {
	digitRunes := slices.Filter([]rune(line), stringutils.IsDigit)
	digits := slices.Map(digitRunes, stringutils.RuneToInt)
	return digits
}

func getCalibrationValue(line string) int {
	digits := getDigits(line)
	firstDigit := digits[0]
	lastDigit := digits[len(digits)-1]
	return firstDigit*10 + lastDigit
}
