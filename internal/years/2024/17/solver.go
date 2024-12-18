package day17

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/intutils"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 17

// init registers the solver for day 17
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 17
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) machine {
	regA, regB, regC := 0, 0, 0
	fmt.Sscanf(lines[0], "Register A: %d", &regA)
	fmt.Sscanf(lines[1], "Register B: %d", &regB)
	fmt.Sscanf(lines[2], "Register C: %d", &regC)
	instructionString := ""
	fmt.Sscanf(lines[4], "Program: %s", &instructionString)
	opCodes := strings.Split(instructionString, ",")
	instructions := slices.Map(opCodes, parseInstruction)
	return newMachine(regA, regB, regC, instructions)
}

type machine struct {
	instructions       []instruction
	instructionPointer int
	regA               int
	regB               int
	regC               int
	output             []int
}

func newMachine(regA, regB, regC int, instructions []instruction) machine {
	return machine{
		instructions:       instructions,
		instructionPointer: 0,
		regA:               regA,
		regB:               regB,
		regC:               regC,
		output:             make([]int, 0),
	}
}

func (m *machine) comboOperand() int {
	operand := m.instructions[m.instructionPointer+1].value()
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return m.regA
	case 5:
		return m.regB
	case 6:
		return m.regC
	case 7:
		panic("Combo operand 7 is reserved")
	default:
		panic("Invalid combo operand")
	}
}

type command func(a int) int

func (m *machine) compile() (int, func(a int) int) {
	variables := map[rune]command{}
	variables['a'] = func(a int) int { return a }
	variables['b'] = func(int) int { return m.regB }
	variables['c'] = func(int) int { return m.regC }
	var program command
	var step int
	for m.instructionPointer < len(m.instructions) && m.instructionPointer >= 0 {
		instruction := m.instructions[m.instructionPointer]
		if instruction.value() == (out{}).value() {
			operand := m.compileComboOperand(variables)
			program = func(a int) int { return operand(a) % 8 }
		}
		if instruction.value() == (adv{}).value() {
			step = intutils.Power(2, m.comboOperand())
		}
		if program == nil {
			r, f := instruction.toFunction(variables, m)
			variables[r] = f
		}
		m.instructionPointer += 2

	}
	return step, program
}

func (m *machine) compileComboOperand(vars map[rune]command) command {
	operand := m.instructions[m.instructionPointer+1].value()
	switch operand {
	case 0, 1, 2, 3:
		return func(int) int {
			return operand
		}
	case 4:
		return vars['a']
	case 5:
		return vars['b']
	case 6:
		return vars['c']
	case 7:
		panic("Combo operand 7 is reserved")
	default:
		panic("Invalid combo operand")
	}
}

func (m *machine) compileLiteralOperand() command {
	value := m.instructions[m.instructionPointer+1].value()
	return func(int) int {
		return value
	}
}

func (m *machine) execute() {
	for m.instructionPointer >= 0 && m.instructionPointer < len(m.instructions) {
		instruction := m.instructions[m.instructionPointer]
		instruction.execute(m)
		m.instructionPointer += 2
	}
}

func (m *machine) literalOperand() int {
	return m.instructions[m.instructionPointer+1].value()
}

type instruction interface {
	execute(m *machine)
	toFunction(variables map[rune]command, m *machine) (rune, command)
	value() int
	fmt.Stringer
}

type adv struct{}

func (adv) execute(m *machine) {
	num := m.regA
	denom := intutils.Power(2, m.comboOperand())
	m.regA = num / denom
}

func (adv) toFunction(vars map[rune]command, m *machine) (rune, command) {
	a := vars['a']
	operand := m.compileComboOperand(vars)
	return 'a', func(init int) int {
		num := a(init)
		denom := intutils.Power(2, operand(init))
		return num / denom
	}
}

func (a adv) value() int {
	return 0
}

func (adv) String() string {
	return "adv"
}

type bxl struct{}

func (bxl) execute(m *machine) {
	m.regB = m.regB ^ m.literalOperand()
}

func (bxl) toFunction(vars map[rune]command, m *machine) (rune, command) {
	b := vars['b']
	operand := m.compileLiteralOperand()
	return 'b', func(init int) int {
		return b(init) ^ operand(init)
	}
}

func (bxl) value() int {
	return 1
}

func (bxl) String() string {
	return "bxl"
}

type bst struct{}

func (bst) execute(m *machine) {
	m.regB = m.comboOperand() % 8
}

func (bst) toFunction(vars map[rune]command, m *machine) (rune, command) {
	operand := m.compileComboOperand(vars)
	return 'b', func(init int) int {
		return operand(init) % 8
	}
}

func (bst) value() int {
	return 2
}

func (bst) String() string {
	return "bst"
}

type jnz struct{}

func (jnz) execute(m *machine) {
	if m.regA == 0 {
		return
	}
	m.instructionPointer = m.literalOperand() - 2
}

func (jnz) toFunction(vars map[rune]command, m *machine) (rune, command) {
	panic("jnz cannot be compiled")
}

func (jnz) value() int {
	return 3
}

func (jnz) String() string {
	return "jnz"
}

type bxc struct{}

func (bxc) execute(m *machine) {
	m.regB = m.regB ^ m.regC
}

func (bxc) toFunction(vars map[rune]command, m *machine) (rune, command) {
	b := vars['b']
	c := vars['c']
	return 'b', func(init int) int {
		return b(init) ^ c(init)
	}
}

func (bxc) value() int {
	return 4
}

func (bxc) String() string {
	return "bxc"
}

type out struct{}

func (out) execute(m *machine) {
	m.output = append(m.output, m.comboOperand()%8)
}

func (out) toFunction(vars map[rune]command, m *machine) (rune, command) {
	panic("out cannot be compiled")
}

func (out) value() int {
	return 5
}

func (out) String() string {
	return "out"
}

type bdv struct{}

func (bdv) execute(m *machine) {
	num := m.regA
	denom := intutils.Power(2, m.comboOperand())
	m.regB = num / denom
}

func (bdv) toFunction(vars map[rune]command, m *machine) (rune, command) {
	a := vars['a']
	operand := m.compileComboOperand(vars)
	return 'b', func(init int) int {
		num := a(init)
		denom := intutils.Power(2, operand(init))
		return num / denom
	}
}

func (bdv) value() int {
	return 6
}

func (bdv) String() string {
	return "bdv"
}

type cdv struct{}

func (cdv) execute(m *machine) {
	num := m.regA
	denom := intutils.Power(2, m.comboOperand())
	m.regC = num / denom
}

func (cdv) toFunction(vars map[rune]command, m *machine) (rune, command) {
	a := vars['a']
	operand := m.compileComboOperand(vars)
	return 'c', func(init int) int {
		num := a(init)
		denom := intutils.Power(2, operand(init))
		return num / denom
	}
}

func (cdv) value() int {
	return 7
}

func (cdv) String() string {
	return "cdv"
}

var instructions = map[int]instruction{
	adv{}.value(): adv{},
	bxl{}.value(): bxl{},
	bst{}.value(): bst{},
	jnz{}.value(): jnz{},
	bxc{}.value(): bxc{},
	out{}.value(): out{},
	bdv{}.value(): bdv{},
	cdv{}.value(): cdv{},
}

func parseInstruction(s string) instruction {
	return instructions[stringutils.Atoi(s)]
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	m := s.parse(lines)
	m.execute()
	return strings.Join(slices.Map(m.output, stringutils.Itoa), ",")
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	m := s.parse(lines)
	opCodes := slices.Map(m.instructions, instruction.value)
	// program2 := func(a int) int {
	// 	return (a % 8) ^ 1 ^ ((a / (intutils.Power(2, ((a % 8) ^ 2)))) % 8)
	// }
	step, program := m.compile()
	possibleA := make([]int, step-1)
	for i := 1; i < step; i++ {
		possibleA[i-1] = i
	}
	for i := len(opCodes) - 1; i >= 0; i-- {
		newPossibleA := []int{}
		for j := 0; j < len(possibleA); j++ {
			a := possibleA[j]
			if program(a) == opCodes[i] {
				base := a * step
				possibles := make([]int, step)
				for k := 0; k < step; k++ {
					possibles[k] = base + k
				}
				newPossibleA = append(newPossibleA, possibles...)
			}
		}
		if len(newPossibleA) == 0 {
			possibleA = slices.Map(possibleA, func(a int) int { return a + 8 })
			i++
			continue
		}
		if i == 0 {
			break
		}
		possibleA = newPossibleA
	}
	for i := 0; i < len(possibleA); i++ {
		if program(possibleA[i]) != opCodes[0] {
			possibleA = slices.RemoveNth(possibleA, i)
			i--
		}
	}
	minA := slices.Min(possibleA)
	return fmt.Sprint(minA)
}
