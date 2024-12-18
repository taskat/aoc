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
	value() int
	fmt.Stringer
}

type adv struct{}

func (adv) execute(m *machine) {
	// fmt.Printf("Adv: regA: %d, op: %d\n", m.regA, m.comboOperand())
	num := m.regA
	denom := intutils.Power(2, m.comboOperand())
	m.regA = num / denom
}

func (a adv) value() int {
	return 0
}

func (adv) String() string {
	return "adv"
}

type bxl struct{}

func (bxl) execute(m *machine) {
	// fmt.Printf("Bxl: regB: %d, op: %d\n", m.regB, m.literalOperand())
	m.regB = m.regB ^ m.literalOperand()
}

func (bxl) value() int {
	return 1
}

func (bxl) String() string {
	return "bxl"
}

type bst struct{}

func (bst) execute(m *machine) {
	// fmt.Printf("Bst: op: %d\n", m.comboOperand())
	m.regB = m.comboOperand() % 8
}

func (bst) value() int {
	return 2
}

func (bst) String() string {
	return "bst"
}

type jnz struct{}

func (jnz) execute(m *machine) {
	// fmt.Printf("Jnz: regA: %d, op: %d\n", m.regA, m.literalOperand())
	if m.regA == 0 {
		return
	}
	m.instructionPointer = m.literalOperand() - 2
}

func (jnz) value() int {
	return 3
}

func (jnz) String() string {
	return "jnz"
}

type bxc struct{}

func (bxc) execute(m *machine) {
	// fmt.Printf("Bxc: regB: %d, regC: %d\n", m.regB, m.regC)
	m.regB = m.regB ^ m.regC
}

func (bxc) value() int {
	return 4
}

func (bxc) String() string {
	return "bxc"
}

type out struct{}

func (out) execute(m *machine) {
	// fmt.Printf("Out: op: %d\n", m.comboOperand())
	m.output = append(m.output, m.comboOperand()%8)
}

func (out) value() int {
	return 5
}

func (out) String() string {
	return "out"
}

type bdv struct{}

func (bdv) execute(m *machine) {
	// fmt.Printf("Bdv: regA: %d, op: %d\n", m.regA, m.comboOperand())
	num := m.regA
	denom := intutils.Power(2, m.comboOperand())
	m.regB = num / denom
}

func (bdv) value() int {
	return 6
}

func (bdv) String() string {
	return "bdv"
}

type cdv struct{}

func (cdv) execute(m *machine) {
	// fmt.Printf("Cdv: regA: %d, op: %d\n", m.regA, m.comboOperand())
	num := m.regA
	denom := intutils.Power(2, m.comboOperand())
	m.regC = num / denom
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
	return ""
}
