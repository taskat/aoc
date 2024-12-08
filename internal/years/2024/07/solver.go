package day07

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/combinatorics"
	"github.com/taskat/aoc/pkg/utils/intutils"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 7

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

// equation represents a poossible equation. It has a result,
// and a list of operands in order, without any operators
type equation struct {
	result   int
	operands []int
}

// parseEquation parses a line into an equation
// The line is expected to be in the format "result: operand1 operand2 ..."
func parseEquation(line string) equation {
	parts := strings.Split(line, ": ")
	var e equation
	e.result = stringutils.Atoi(parts[0])
	operands := strings.Split(parts[1], " ")
	e.operands = slices.Map(operands, stringutils.Atoi)
	return e
}

// canProduce returns true if the equation can be evaluated to the result, using
// the given operators. It uses the CartesianProduct of the operators to find
// all possible orderings of the operators
func (e equation) canProduce(possibleOpeartors []operator) bool {
	orderings := cartesianWithCache(possibleOpeartors, len(e.operands)-1)
	return slices.Any(orderings, e.isPossible)
}

var cartesianCache = make(map[int][][]operator)

func cartesianWithCache(possibleOpeartors []operator, length int) [][]operator {
	if v, ok := cartesianCache[length]; ok {
		return v
	}
	orderings := combinatorics.CartesianProduct(possibleOpeartors, length)
	cartesianCache[length] = orderings
	return orderings
}

// isPossible returns true if the equation can be evaluated to the result
// using the given operators
func (e equation) isPossible(operators []operator) bool {
	f := func(left, right, index int) int {
		return operators[index].eval(left, right)
	}
	result := slices.Reduce_i(e.operands[1:], f, e.operands[0])
	return result == e.result
}

// operator is an interface that defines an operator that can be used in an equation
type operator interface {
	eval(left, right int) int
}

// sum is an operator that adds two numbers
type sum struct{}

// eval adds two numbers
func (s sum) eval(left, right int) int {
	return left + right
}

// String returns the string representation of the operator
func (s sum) String() string {
	return "+"
}

// product is an operator that multiplies two numbers
type product struct{}

// eval multiplies two numbers
func (p product) eval(left, right int) int {
	return left * right
}

// String returns the string representation of the operator
func (p product) String() string {
	return "*"
}

// concat is an operator that concatenates two numbers
type concat struct{}

// eval concatenates two numbers
func (c concat) eval(left, right int) int {
	offset := intutils.Power(10, intutils.Length(right))
	return left*offset + right
}

// String returns the string representation of the operator
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
