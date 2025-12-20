package day10

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 10

// init registers the solver for day 10
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 10
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []machine {
	return slices.Map(lines, parseMachine)

}

// machine represents the description of the machine
type machine struct {
	indicators lights
	wirings    []lights
}

// parseMachine parses the machine from input line
func parseMachine(line string) machine {
	line = line[1:]
	parts := strings.Split(line, "] (")
	indicatorLights := parseIndicatorLights(parts[0])
	parts = strings.Split(parts[1], ") {")
	wiringLights := parseWiringLights(parts[0], len(indicatorLights))
	return machine{
		indicators: indicatorLights,
		wirings:    wiringLights,
	}
}

// minimalButtonPresses calculates the minimal button presses needed to match the indicator lights
func (m machine) minimalButtonPresses() int {
	start := make(lights, len(m.indicators))
	minPresses, _ := m.minimalButtonPressesRecursion(start, 0)
	return minPresses
}

// minimalButtonPressesRecursion calculates the minimal button presses needed to match the indicator lights
// starting from the current lights configuration and considering wirings from firstWiring index.
// It returns the minimal number of presses and a boolean indicating if a solution was found.
func (m machine) minimalButtonPressesRecursion(current lights, firstWiring int) (int, bool) {
	if current.matchToggle(m.indicators) {
		return 0, true
	}
	minimalPresses := len(m.wirings) - firstWiring + 1
	for i := firstWiring; i < len(m.wirings); i++ {
		current.toggle(m.wirings[i])
		presses, found := m.minimalButtonPressesRecursion(current, i+1)
		currentPresses := presses + 1
		if found && currentPresses < minimalPresses {
			minimalPresses = currentPresses
		}
		current.toggle(m.wirings[i]) // backtrack
	}
	if minimalPresses == len(m.wirings)-firstWiring+1 {
		return 0, false
	}
	return minimalPresses, true
}

// lights represents the indicator lights
type lights []int

// matchToggle checks if two lights configurations match
// considering the number of toggles. E.g. if both have the same lights on/off
func (l lights) matchToggle(other lights) bool {
	if len(l) != len(other) {
		return false
	}
	for i := range l {
		if l[i]%2 != other[i]%2 {
			return false
		}
	}
	return true
}

// toggle toggles the lights based on another lights configuration
func (l lights) toggle(other lights) {
	for i, toggle := range other {
		l[i] += toggle
	}
}

// parseIndictorLights parses the indicator lights from input lines
// The indicator lights are represented as a string of '.' and '#' characters
func parseIndicatorLights(line string) lights {
	lights := make(lights, len(line))
	for i, ch := range line {
		if ch == '#' {
			lights[i] = 1
		} else {
			lights[i] = 0
		}
	}
	return lights
}

// parseWiringLights parses the wiring lights from input lines
// The wiring consist of multiple buttons, each button separated by ") (".
func parseWiringLights(line string, numberOfLigths int) []lights {
	buttons := strings.Split(line, ") (")
	lights := slices.Map(buttons, func(button string) lights { return parseButton(button, numberOfLigths) })
	return lights
}

// parseButton parses a single button's lights
func parseButton(button string, numberOfLights int) lights {
	lights := make(lights, numberOfLights)
	parts := strings.SplitSeq(button, ",")
	for part := range parts {
		index := stringutils.Atoi(part)
		lights[index] = 1
	}
	return lights
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	machines := s.parse(lines)
	minimalPresses := slices.Map(machines, machine.minimalButtonPresses)
	return fmt.Sprintf("%d", slices.Sum(minimalPresses))
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
