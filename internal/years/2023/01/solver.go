package day01

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2023/days"
	"github.com/taskat/aoc/pkg/utils/maps"
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
	getCalibrationValue := func(line string) int {
		return getCalibrationValue(line, digits)
	}
	calibrationValues := slices.Map(s.lines, getCalibrationValue)
	sum := slices.Sum(calibrationValues)
	return fmt.Sprintf("%d", sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2() string {
	digits := maps.Merge(digits, extraDigits)
	getCalibrationValue := func(line string) int {
		return getCalibrationValue(line, digits)
	}
	calibrationValues := slices.Map(s.lines, getCalibrationValue)
	sum := slices.Sum(calibrationValues)
	return fmt.Sprintf("%d", sum)
}

func getDigits(line string, possibleDigits map[string]int) []int {
	digits := make([]int, 0)
	for i := range line {
		for extraDigit, value := range possibleDigits {
			if strings.HasPrefix(line[i:], extraDigit) {
				digits = append(digits, value)
				break
			}
		}
	}
	return digits
}

func getCalibrationValue(line string, possibleDigits map[string]int) int {
	digits := getDigits(line, possibleDigits)
	firstDigit := digits[0]
	lastDigit := digits[len(digits)-1]
	return firstDigit*10 + lastDigit
}

var digits = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var extraDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
