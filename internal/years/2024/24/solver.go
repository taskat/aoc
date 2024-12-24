package day24

import (
	"fmt"
	"sort"
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
	getInput1() wire
	getInput2() wire
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

func (g *baseGate) execute(wires map[wire]bool) bool {
	panic("Not implemented")
}

func (g *baseGate) getInput1() wire {
	return g.input1
}

func (g *baseGate) getInput2() wire {
	return g.input2
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

func findWire[T gate](gates []gate, input1, input2 wire) wire {
	for _, g := range gates {
		gate, ok := g.(T)
		if !ok {
			continue
		}
		if (gate.getInput1() == input1 && gate.getInput2() == input2) || (gate.getInput1() == input2 && gate.getInput2() == input1) {
			return gate.getOutput()
		}
	}
	return ""
}

func findGateByInput[T gate](gates []gate, input wire) (T, bool) {
	for _, g := range gates {
		gate, ok := g.(T)
		if !ok {
			continue
		}
		if gate.getInput1() == input || gate.getInput2() == input {
			return gate, true
		}
	}
	var zero T
	return zero, false
}

func finGateByOutput[T gate](gates []gate, output wire) (T, bool) {
	for _, g := range gates {
		gate, ok := g.(T)
		if !ok {
			continue
		}
		if g.getOutput() == output {
			return gate, true
		}
	}
	var zero T
	return zero, false
}

func check0Bit(gates []gate) (wire, []wire) {
	carryWire := findWire[*andGate](gates, "x00", "y00")
	resultWire := findWire[*xorGate](gates, "x00", "y00")
	if resultWire != "z00" {
		return carryWire, []wire{resultWire, wire("z00")}
	}
	return carryWire, nil
}

func checkNthBit(gates []gate, n int, carry wire) (wire, []wire) {
	wrongs := make([]wire, 0)
	xWire := wire(fmt.Sprintf("x%02d", n))
	yWire := wire(fmt.Sprintf("y%02d", n))
	zWire := wire(fmt.Sprintf("z%02d", n))
	bitResult := findWire[*xorGate](gates, xWire, yWire)
	resultWire := findWire[*xorGate](gates, bitResult, carry)
	zGate, ok := finGateByOutput[*xorGate](gates, zWire)
	if !ok {
		wrongs = append(wrongs, zWire, resultWire)
	} else {
		switch resultWire {
		case "":
			// carry or bitrsult is wrong
			if zGate.getInput1() == carry {
				wrongs = append(wrongs, zGate.getInput2(), bitResult)
				bitResult = zGate.getInput2()
			} else if zGate.getInput2() == carry {
				wrongs = append(wrongs, zGate.getInput1(), bitResult)
				bitResult = zGate.getInput1()
			} else if zGate.getInput1() == bitResult {
				wrongs = append(wrongs, zGate.getInput2(), carry)
				carry = zGate.getInput2()
			} else if zGate.getInput2() == bitResult {
				wrongs = append(wrongs, zGate.getInput1(), carry)
				carry = zGate.getInput1()
			} else {
				// both wrong
				wrongs = append(wrongs, zGate.getInput1(), zGate.getInput2(), carry, bitResult)
			}
		case zWire:
			// carry and bitresult is ok
		default:
			// output is wrong
			wrongs = append(wrongs, resultWire, zWire)
		}
	}
	resultCarry := findWire[*andGate](gates, carry, bitResult)
	bitCarry := findWire[*xorGate](gates, xWire, yWire)
	newCarry := findWire[*orGate](gates, bitCarry, resultCarry)
	if newCarry != "" {
		return newCarry, wrongs
	}
	carryGate, ok := findGateByInput[*andGate](gates, bitCarry)
	if ok {
		if carryGate.getInput1() == bitCarry {
			wrongs = append(wrongs, carryGate.getInput2(), resultCarry)
		} else {
			wrongs = append(wrongs, carryGate.getInput1(), resultCarry)
		}
		return carryGate.output, wrongs
	}
	carryGate, ok = findGateByInput[*andGate](gates, resultCarry)
	if ok {
		if carryGate.getInput1() == resultCarry {
			wrongs = append(wrongs, carryGate.getInput2(), bitCarry)
		} else {
			wrongs = append(wrongs, carryGate.getInput1(), bitCarry)
		}
		return carryGate.output, wrongs
	}
	fmt.Println(n)
	panic("unrecoverable")
}

func checkLastBit(n int, carry wire) []wire {
	zWire := wire(fmt.Sprintf("z%02d", n))
	if carry != zWire {
		return []wire{zWire, carry}
	}
	return nil
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	_, gates := s.parse(lines)
	zWires := slices.Filter(gates, func(g gate) bool { return strings.HasPrefix(string(g.getOutput()), "z") })
	wireNames := slices.Map(zWires, func(g gate) string { return string(g.getOutput()) })
	sort.Strings(wireNames)
	maxZ := wireNames[len(wireNames)-1]
	maxBit := stringutils.Atoi(maxZ[1:])
	carry, wrongs := check0Bit(gates)
	for i := 1; i < maxBit; i++ {
		var newWrongs []wire
		carry, newWrongs = checkNthBit(gates, i, carry)
		wrongs = append(wrongs, newWrongs...)
	}
	newWrongs := checkLastBit(maxBit, carry)
	wrongs = append(wrongs, newWrongs...)
	fmt.Println(wrongs)
	return ""
}
