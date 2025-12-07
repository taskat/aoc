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
func (s *Solver) parse(lines []string, part int) []problem {
	if part == 1 {
		return s.parseVerticalProblems(lines)
	}
	if part == 2 {
		return s.parseCephalopodProblems(lines)
	}
	panic("unknown part")
}

// parseVerticalProblems parses the input lines, as if they were vertical problems
func (s *Solver) parseVerticalProblems(lines []string) []problem {
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
	s.parseAndAddoperators(lines[len(lines)-1], problems)
	return problems
}

// parseLine parses a single line into an array, without the unnecessary spaces
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

// parseAndAddoperators parses the operators from the line and adds them to the problems
func (s *Solver) parseAndAddoperators(line string, problems []problem) {
	operators := s.parseLine(line)
	for i, operator := range operators {
		switch operator {
		case "+":
			problems[i].operator = slices.Sum
		case "*":
			problems[i].operator = slices.Product
		}
	}
}

// parseCephalopodProblems parses the input lines, as if they were cephalopod problems
func (s *Solver) parseCephalopodProblems(lines []string) []problem {
	problems := []problem{}
	newProblem := problem{}
	for j := range lines[0] {
		digits := []string{}
		for i := 0; i < len(lines)-1; i++ {
			if stringutils.IsDigit(rune(lines[i][j])) {
				digits = append(digits, string(lines[i][j]))
			}
		}
		if len(digits) > 0 {
			newNumber := stringutils.Atoi(strings.Join(digits, ""))
			newProblem.numbers = append(newProblem.numbers, newNumber)
			digits = []string{}
		} else {
			problems = append(problems, newProblem)
			newProblem = problem{}
		}
	}
	problems = append(problems, newProblem)
	s.parseAndAddoperators(lines[len(lines)-1], problems)
	return problems
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
	problems := s.parse(lines, 1)
	results := slices.Map(problems, problem.solve)
	sum := slices.Sum(results)
	return fmt.Sprint(sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	problems := s.parse(lines, 2)
	results := slices.Map(problems, problem.solve)
	sum := slices.Sum(results)
	return fmt.Sprint(sum)
}
