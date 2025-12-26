package day10

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/math"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"

	"github.com/draffensperger/golp"
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
	joltage    lights
}

// parseMachine parses the machine from input line
func parseMachine(line string) machine {
	line = line[1 : len(line)-1]
	parts := strings.Split(line, "] (")
	indicatorLights := parseIndicatorLights(parts[0])
	parts = strings.Split(parts[1], ") {")
	wiringLights := parseWiringLights(parts[0], len(indicatorLights))
	joltageRequirement := parseJoltage(parts[1])
	return machine{
		indicators: indicatorLights,
		wirings:    wiringLights,
		joltage:    joltageRequirement,
	}
}

// minimalButtonPressesToConfigureLights calculates the minimal button presses needed to match the indicator lights
func (m machine) minimalButtonPressesToConfigureLights() int {
	start := make(lights, len(m.indicators))
	minPresses, _ := m.minimalButtonPressesToConfigureLights_recursion(start, 0)
	return minPresses
}

// minimalButtonPressesToConfigureLights_recursion calculates the minimal button presses needed to match the indicator lights
// starting from the current lights configuration and considering wirings from firstWiring index.
// It returns the minimal number of presses and a boolean indicating if a solution was found.
func (m machine) minimalButtonPressesToConfigureLights_recursion(current lights, firstWiring int) (int, bool) {
	if current.matchToggle(m.indicators) {
		return 0, true
	}
	minimalPresses := len(m.wirings) - firstWiring + 1
	for i := firstWiring; i < len(m.wirings); i++ {
		current.toggle(m.wirings[i])
		presses, found := m.minimalButtonPressesToConfigureLights_recursion(current, i+1)
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

// minimalButtonPressesToMatchJoltage calculates the minimal button presses needed to match the joltage requirement
func (m machine) minimalButtonPressesToMatchJoltage() int {
	lp := golp.NewLP(len(m.indicators), len(m.wirings))
	for i, expectedJoltage := range m.joltage {
		coefficients := make([]float64, len(m.wirings))
		for j, wiring := range m.wirings {
			coefficients[j] = float64(wiring[i])
		}
		lp.AddConstraint(coefficients, golp.EQ, float64(expectedJoltage))
	}
	for i := range m.wirings {
		lp.SetInt(i, true)
	}
	objective := make([]float64, len(m.wirings))
	for i := range objective {
		objective[i] = 1
	}
	lp.SetObjFn(objective)
	solution := lp.Solve()
	if solution != golp.OPTIMAL {
		panic("No optimal solution found")
	}
	return math.Round(lp.Objective())
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

// parseJoltage parse the joltage requirement from input string
func parseJoltage(line string) lights {
	parts := strings.Split(line, ",")
	lights := make(lights, len(parts))
	for i, part := range parts {
		lights[i] = stringutils.Atoi(part)
	}
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
	minimalPresses := slices.Map(machines, machine.minimalButtonPressesToConfigureLights)
	return fmt.Sprintf("%d", slices.Sum(minimalPresses))
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	machines := s.parse(lines)
	minimalPresses := slices.Map(machines, machine.minimalButtonPressesToMatchJoltage)
	return fmt.Sprintf("%d", slices.Sum(minimalPresses))
}
