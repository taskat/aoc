package day24

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/maps"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 24

// init registers the solver for day 24
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 24
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) (map[wire]bool, []gate) {
	idx := slices.FindIndex(lines, stringutils.IsEmpty)
	wires := make(map[wire]bool, idx)
	for _, line := range lines[:idx] {
		parts := strings.Split(line, ": ")
		wires[wire(parts[0])] = stringutils.StringToBool(parts[1])
	}
	gates := make([]gate, len(lines)-idx-1)
	for i, line := range lines[idx+1:] {
		gates[i] = parseGate(line)
	}
	return wires, gates
}

type wire string

type gate interface {
	canExecute(wires map[wire]bool) bool
	execute(wires map[wire]bool) bool
	getOutput() wire
}

func parseGate(line string) gate {
	input1, input2, op, output := "", "", "", ""
	_, err := fmt.Sscanf(line, "%s %s %s -> %s", &input1, &op, &input2, &output)
	if err != nil {
		panic(err)
	}
	base := baseGate{
		input1: wire(input1),
		input2: wire(input2),
		output: wire(output),
	}
	switch op {
	case "AND":
		return &andGate{base}
	case "OR":
		return &orGate{base}
	case "XOR":
		return &xorGate{base}
	}
	panic("Invalid gate: " + op)
}

type baseGate struct {
	input1 wire
	input2 wire
	output wire
}

func (g *baseGate) canExecute(wires map[wire]bool) bool {
	_, ok1 := wires[g.input1]
	_, ok2 := wires[g.input2]
	return ok1 && ok2
}

func (g *baseGate) getOutput() wire {
	return g.output
}

type andGate struct{ baseGate }

func (g *andGate) execute(wires map[wire]bool) bool {
	return wires[g.input1] && wires[g.input2]
}

type orGate struct{ baseGate }

func (g *orGate) execute(wires map[wire]bool) bool {
	return wires[g.input1] || wires[g.input2]
}

type xorGate struct{ baseGate }

func (g *xorGate) execute(wires map[wire]bool) bool {
	return wires[g.input1] != wires[g.input2]
}

func executeAll(wires map[wire]bool, gates []gate) map[wire]bool {
	for {
		executedAll := true
		for _, g := range gates {
			if g.canExecute(wires) {
				wires[g.getOutput()] = g.execute(wires)
			} else {
				executedAll = false
			}
		}
		if executedAll {
			break
		}
	}
	return wires
}

func bitsToInt(bits []bool) int {
	num := 0
	for _, bit := range bits {
		if bit {
			num += 1
		}
		num <<= 1
	}
	return num >> 1
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	wires, gates := s.parse(lines)
	wires = executeAll(wires, gates)
	outputs := maps.Filter(wires, func(w wire, v bool) bool {
		return strings.HasPrefix(string(w), "z")
	})
	outputWires := maps.Keys(outputs)
	slices.Sort(outputWires, func(a, b wire) bool {
		return a > b
	})
	bits := slices.Map(outputWires, func(w wire) bool {
		return outputs[w]
	})
	number := bitsToInt(bits)
	return fmt.Sprint(number)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
