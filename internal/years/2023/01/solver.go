package day01

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2023/days"
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
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) {}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	getCalibrationValue := func(line string) int {
		return getCalibrationValue(line, digits)
	}
	calibrationValues := slices.Map(lines, getCalibrationValue)
	sum := slices.Sum(calibrationValues)
	return fmt.Sprintf("%d", sum)
}

// getCalibrationValue returns the calibration value for a line
// based on the possible digits
func getCalibrationValue(line string, possibleDigits map[string]int) int {
	digits := getDigits(line, possibleDigits)
	firstDigit := digits[0]
	lastDigit := digits[len(digits)-1]
	return firstDigit*10 + lastDigit
}

// getDigits returns the digits of a line based on the possible digits.
// It returns all the digits, even overlapping ones (for part 2).
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

// digits is a map of possible digits made of digits
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

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
