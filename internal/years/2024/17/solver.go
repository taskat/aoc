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
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 17
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

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

// machine represents a virtual machine that can run the given instructions
// It contains the instructions, the instructionPointer to keep track which
// instruction it currently executes, the values in the registers, and also
// the outputs.
type machine struct {
	instructions       []instruction
	instructionPointer int
	regA               int
	regB               int
	regC               int
	output             []int
}

// newMachine creates a new machines and initializes it with the given values.
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

// comboOperand return the next instruction's operand's value as a combo operand
func (m *machine) comboOperand() int {
	operand := m.instructions[m.instructionPointer+1].opcode()
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

// command represents a compiled program, that takes a parameter as the default value of regsiter A
type command func(a int) int

// compile compiles the machine's instruction into a function call that represents the looped logic
// in the instructions. It returns the step size (from adv instruction) and the compiled program.
// The step size is the value that register A should be moduloed after each loop.
func (m *machine) compile() (int, func(a int) int) {
	variables := map[rune]command{}
	variables['a'] = func(a int) int { return a }
	variables['b'] = func(int) int { return m.regB }
	variables['c'] = func(int) int { return m.regC }
	var program command
	var step int
	for m.instructionPointer < len(m.instructions) && m.instructionPointer >= 0 {
		instruction := m.instructions[m.instructionPointer]
		if instruction.opcode() == (out{}).opcode() {
			operand := m.compileComboOperand(variables)
			program = func(a int) int { return operand(a) % 8 }
		}
		if instruction.opcode() == (adv{}).opcode() {
			step = intutils.Power(2, m.comboOperand())
		}
		if program == nil {
			r, f := instruction.compile(variables, m)
			variables[r] = f
		}
		m.instructionPointer += 2

	}
	return step, program
}

// compileComboOperand returns a command that returns the combo operand's value
func (m *machine) compileComboOperand(vars map[rune]command) command {
	operand := m.instructions[m.instructionPointer+1].opcode()
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
	value := m.instructions[m.instructionPointer+1].opcode()
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
	return m.instructions[m.instructionPointer+1].opcode()
}

// instruction represents a single type of instruction that the machine can execute
type instruction interface {
	// execute executes the instruction itslef using the current state of the machine as parameters
	execute(m *machine)
	// compile compiles the instruction into a command that can be invoked as a function call
	// It also returns a rune, that can be used to represent, if the compiled command's output
	// should be saved in register
	compile(variables map[rune]command, m *machine) (rune, command)
	// opcode returns the opcode of the instruction
	opcode() int
	fmt.Stringer
}

// adv is a division instruction. The numerator is the value in the A register. The denominator
// is found by raising 2 to the power of the instruction's combo operand. The result of the
// division operation is truncated to an integer and then written to the A register.
type adv struct{}

func (adv) execute(m *machine) {
	num := m.regA
	denom := intutils.Power(2, m.comboOperand())
	m.regA = num / denom
}

func (adv) compile(vars map[rune]command, m *machine) (rune, command) {
	a := vars['a']
	operand := m.compileComboOperand(vars)
	return 'a', func(init int) int {
		num := a(init)
		denom := intutils.Power(2, operand(init))
		return num / denom
	}
}

func (a adv) opcode() int {
	return 0
}

func (adv) String() string {
	return "adv"
}

// bxl is a bitwise xor of register B and the instruction's literal operand,
// then stores the result in register B.
type bxl struct{}

func (bxl) execute(m *machine) {
	m.regB = m.regB ^ m.literalOperand()
}

func (bxl) compile(vars map[rune]command, m *machine) (rune, command) {
	b := vars['b']
	operand := m.compileLiteralOperand()
	return 'b', func(init int) int {
		return b(init) ^ operand(init)
	}
}

func (bxl) opcode() int {
	return 1
}

func (bxl) String() string {
	return "bxl"
}

// bst is a modulo instruction that calculates the value of its combo operand
// modulo 8 (thereby keeping only its lowest 3 bits), then writes that value to the B register.
type bst struct{}

func (bst) execute(m *machine) {
	m.regB = m.comboOperand() % 8
}

func (bst) compile(vars map[rune]command, m *machine) (rune, command) {
	operand := m.compileComboOperand(vars)
	return 'b', func(init int) int {
		return operand(init) % 8
	}
}

func (bst) opcode() int {
	return 2
}

func (bst) String() string {
	return "bst"
}

// jnz is a jump instruction. It does nothing if regster A is 0. If it is not 0, it jumps by setting
// the instruciton pointer to the value of the literal operand (that way, the instruction pointer
// does not get increased by 2)
type jnz struct{}

func (jnz) execute(m *machine) {
	if m.regA == 0 {
		return
	}
	m.instructionPointer = m.literalOperand() - 2
}

func (jnz) compile(vars map[rune]command, m *machine) (rune, command) {
	panic("jnz cannot be compiled")
}

func (jnz) opcode() int {
	return 3
}

func (jnz) String() string {
	return "jnz"
}

// bxc is a bitwise XOR of register B and register C, then stores the result in register B.
type bxc struct{}

func (bxc) execute(m *machine) {
	m.regB = m.regB ^ m.regC
}

func (bxc) compile(vars map[rune]command, m *machine) (rune, command) {
	b := vars['b']
	c := vars['c']
	return 'b', func(init int) int {
		return b(init) ^ c(init)
	}
}

func (bxc) opcode() int {
	return 4
}

func (bxc) String() string {
	return "bxc"
}

// out outputs its combo operand's value modulo 8
type out struct{}

func (out) execute(m *machine) {
	m.output = append(m.output, m.comboOperand()%8)
}

func (out) compile(vars map[rune]command, m *machine) (rune, command) {
	panic("out cannot be compiled")
}

func (out) opcode() int {
	return 5
}

func (out) String() string {
	return "out"
}

// bdv is a division instruction. The numerator is the value in the A register. The denominator
// is found by raising 2 to the power of the instruction's combo operand. The result of the
// division operation is truncated to an integer and then written to the B register.
type bdv struct{}

func (bdv) execute(m *machine) {
	num := m.regA
	denom := intutils.Power(2, m.comboOperand())
	m.regB = num / denom
}

func (bdv) compile(vars map[rune]command, m *machine) (rune, command) {
	a := vars['a']
	operand := m.compileComboOperand(vars)
	return 'b', func(init int) int {
		num := a(init)
		denom := intutils.Power(2, operand(init))
		return num / denom
	}
}

func (bdv) opcode() int {
	return 6
}

func (bdv) String() string {
	return "bdv"
}

// adv is a division instruction. The numerator is the value in the A register. The denominator
// is found by raising 2 to the power of the instruction's combo operand. The result of the
// division operation is truncated to an integer and then written to the C register.
type cdv struct{}

func (cdv) execute(m *machine) {
	num := m.regA
	denom := intutils.Power(2, m.comboOperand())
	m.regC = num / denom
}

func (cdv) compile(vars map[rune]command, m *machine) (rune, command) {
	a := vars['a']
	operand := m.compileComboOperand(vars)
	return 'c', func(init int) int {
		num := a(init)
		denom := intutils.Power(2, operand(init))
		return num / denom
	}
}

func (cdv) opcode() int {
	return 7
}

func (cdv) String() string {
	return "cdv"
}

// instructions contains all the possible instructions, mapped to their opcode
var instructions = map[int]instruction{
	adv{}.opcode(): adv{},
	bxl{}.opcode(): bxl{},
	bst{}.opcode(): bst{},
	jnz{}.opcode(): jnz{},
	bxc{}.opcode(): bxc{},
	out{}.opcode(): out{},
	bdv{}.opcode(): bdv{},
	cdv{}.opcode(): cdv{},
}

// parseInstruction returns the parsed instruction
func parseInstruction(s string) instruction {
	return instructions[stringutils.Atoi(s)]
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	m := s.parse(lines)
	m.execute()
	return strings.Join(slices.Map(m.output, stringutils.Itoa), ",")
}

// findInputs finds the possible inputs A, that can be used as initial value for register A
// to output the given output value. The search is done between base and base+step.
func findInputs(program command, output, base, step int) []int {
	inputs := make([]int, 0, step)
	if program(base) == output {
		nextBase := base * step
		for k := range step {
			inputs = append(inputs, nextBase+k)
		}
	}
	return inputs
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	m := s.parse(lines)
	desiredOutput := slices.Map(m.instructions, instruction.opcode)
	// program2 := func(a int) int {
	// 	return (a % 8) ^ 1 ^ ((a / (intutils.Power(2, ((a % 8) ^ 2)))) % 8)
	// }
	step, program := m.compile()
	bases := make([]int, step-1)
	for i := 1; i < step; i++ {
		bases[i-1] = i
	}
	for i := len(desiredOutput) - 1; i >= 0; i-- {
		newBases := []int{}
		for j := 0; j < len(bases); j++ {
			a := bases[j]
			inputs := findInputs(program, desiredOutput[i], a, step)
			newBases = append(newBases, inputs...)
		}
		if len(newBases) == 0 {
			bases = slices.Map(bases, func(a int) int { return a + step })
			i++
			continue
		}
		if i != 0 {
			bases = newBases
		}
	}
	for i := 0; i < len(bases); i++ {
		if program(bases[i]) != desiredOutput[0] {
			bases = slices.RemoveNth(bases, i)
			i--
		}
	}
	minA := slices.Min(bases)
	return fmt.Sprint(minA)
}
