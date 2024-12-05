package day05

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 05

// init registers the solver for day 05
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 05
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) ([]rule, []manual) {
	index := slices.FindIndex(lines, func(s string) bool { return s == "" })
	rules := slices.Map(lines[:index], parseRule)
	manuals := slices.Map(lines[index+1:], parseManual)
	return rules, manuals
}

type rule struct {
	before int
	after  int
}

func parseRule(s string) rule {
	before, after := 0, 0
	_, _ = fmt.Sscanf(s, "%d|%d", &before, &after)
	return rule{before: before, after: after}
}

type manual []int

func parseManual(s string) manual {
	pages := strings.Split(s, ",")
	return slices.Map(pages, stringutils.Atoi)
}

func (m manual) fixOrder(rules []rule) {
	if m.isCorrect(rules) {
		return
	}
	for i, page := range m {
		for _, r := range rules {
			if r.after != page {
				continue
			}
			for j := i + 1; j < len(m); j++ {
				if m[j] == r.before {
					m[i], m[j] = m[j], m[i]
					m.fixOrder(rules)
					return
				}
			}
		}
	}
}

func (m manual) isCorrect(rules []rule) bool {
	for i, page := range m {
		for _, r := range rules {
			if r.after != page {
				continue
			}
			for j := i + 1; j < len(m); j++ {
				if m[j] == r.before {
					return false
				}
			}
		}
	}
	return true
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	rules, manuals := s.parse(lines)
	manuals = slices.Filter(manuals, func(m manual) bool { return m.isCorrect(rules) })
	middlePages := slices.Map(manuals, func(m manual) int { return m[len(m)/2] })
	sum := slices.Sum(middlePages)
	return fmt.Sprintf("%d", sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	rules, manuals := s.parse(lines)
	manuals = slices.Filter(manuals, func(m manual) bool { return !m.isCorrect(rules) })
	slices.ForEach(manuals, func(m manual) { m.fixOrder(rules) })
	middlePages := slices.Map(manuals, func(m manual) int { return m[len(m)/2] })
	sum := slices.Sum(middlePages)
	return fmt.Sprintf("%d", sum)
}
