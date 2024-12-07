package day07

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/intutils"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 07

// init registers the solver for day 07
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 07
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []equation {
	return slices.Map(lines, parseEquation)
}

type equation struct {
	result   int
	operands []int
}

func parseEquation(line string) equation {
	parts := strings.Split(line, ": ")
	var e equation
	e.result = stringutils.Atoi(parts[0])
	operands := strings.Split(parts[1], " ")
	e.operands = slices.Map(operands, stringutils.Atoi)
	return e
}

func (e equation) canProduce(possuibleOpeartors []operator) bool {
	orderings := cartesianProduct(possuibleOpeartors, len(e.operands)-1)
	for _, ordering := range orderings {
		if e.isPossible(ordering) {
			fmt.Println("found", e, ordering)
			return true
		}
	}
	fmt.Println("not found", e)
	return false
}

func (e equation) isPossible(operators []operator) bool {
	result := e.operands[0]
	for i, operator := range operators {
		result = operator.eval(result, e.operands[i+1])
	}
	return result == e.result
}

func cartesianProduct(element []operator, length int) [][]operator {
	if length == 0 {
		return [][]operator{{}}
	}
	result := make([][]operator, 0, intutils.Power(len(element), length))
	for _, e := range element {
		rest := cartesianProduct(element, length-1)
		for _, r := range rest {
			result = append(result, append([]operator{e}, r...))
		}
	}
	return result
}

type operator interface {
	eval(left, right int) int
}

type sum struct{}

func (s sum) eval(left, right int) int {
	return left + right
}

func (s sum) String() string {
	return "+"
}

type product struct{}

func (p product) eval(left, right int) int {
	return left * right
}

func (p product) String() string {
	return "*"
}

type concat struct{}

func (c concat) eval(left, right int) int {
	offset := intutils.Power(10, intutils.Length(right))
	return left*offset + right
}

func (c concat) String() string {
	return "||"
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	equations := s.parse(lines)
	operators := []operator{sum{}, product{}}
	possibleEquations := slices.Filter(equations, func(e equation) bool { return e.canProduce(operators) })
	sum := slices.Sum(slices.Map(possibleEquations, func(e equation) int { return e.result }))
	return fmt.Sprint(sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	equations := s.parse(lines)
	operators := []operator{sum{}, product{}, concat{}}
	possibleEquations := slices.Filter(equations, func(e equation) bool { return e.canProduce(operators) })
	sum := slices.Sum(slices.Map(possibleEquations, func(e equation) int { return e.result }))
	return fmt.Sprint(sum)
}
