package day06

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 6

// init registers the solver for day 06
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 06
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []problem {
	problems := []problem{}
	for i, line := range lines {
		if i == len(lines)-1 {
			break
		}
		parts := s.parseLine(line)
		if len(problems) == 0 {
			problems = make([]problem, len(parts))
		}
		for j, part := range parts {
			problems[j].numbers = append(problems[j].numbers, stringutils.Atoi(part))
		}
	}
	operators := s.parseLine(lines[len(lines)-1])
	for i, operator := range operators {
		switch operator {
		case "+":
			problems[i].operator = slices.Sum
		case "*":
			problems[i].operator = slices.Product
		}
	}
	return problems
}

// parseLine parses a single line into ana array, without the unnecessary spaces
func (s *Solver) parseLine(line string) []string {
	parts := strings.Split(line, " ")
	for i := 0; i < len(parts); i++ {
		part := parts[i]
		part = strings.Trim(part, " ")
		if part == "" {
			parts = append(parts[:i], parts[i+1:]...)
			i--
		}
	}
	return parts
}

type problem struct {
	numbers  []int
	operator func([]int) int
}

// solve calculates the result of the problem
func (p problem) solve() int {
	return p.operator(p.numbers)
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	problems := s.parse(lines)
	results := slices.Map(problems, problem.solve)
	sum := slices.Sum(results)
	return fmt.Sprint(sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
