package day05

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 5

// init registers the solver for day 05
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 05
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) ([]rule, []manual) {
	index := slices.FindIndex(lines, func(s string) bool { return s == "" })
	rules := slices.Map(lines[:index], parseRule)
	manuals := slices.Map(lines[index+1:], parseManual)
	return rules, manuals
}

// rule represents a page ordering rule. The page with number `before` must be
// before the page with number `after`
type rule struct {
	before int
	after  int
}

// parseRule parses a rule from a string with the format "before|after"
func parseRule(s string) rule {
	before, after := 0, 0
	_, _ = fmt.Sscanf(s, "%d|%d", &before, &after)
	return rule{before: before, after: after}
}

// manual represents a manual with the order of the pages
type manual []int

// parseManual parses a manual from a string with the format "page1,page2,..."
func parseManual(s string) manual {
	pages := strings.Split(s, ",")
	return slices.Map(pages, stringutils.Atoi)
}

// findProblematicPages returns the indices of two page numbers that are in
// the wrong order according to the given rules. If all pages are in the correct
// order, it returns -1, -1
func (m manual) findProblematicPages(rules []rule) (int, int) {
	for i, page := range m {
		for _, r := range rules {
			if r.after != page {
				continue
			}
			for j := i + 1; j < len(m); j++ {
				if m[j] == r.before {
					return i, j
				}
			}
		}
	}
	return -1, -1
}

// fixOrder tries to fix the order of the pages in the manual according to the
// given rules. It does not return anything, it just modifies the manual in
// place. If the manual is already correct, it does nothing.
func (m manual) fixOrder(rules []rule) {
	if m.isCorrect(rules) {
		return
	}
	before, after := m.findProblematicPages(rules)
	slices.Swap(m, before, after)
	m.fixOrder(rules)
	return
}

// isCorrect returns true if the manual is in the correct order according to the
// given rules
func (m manual) isCorrect(rules []rule) bool {
	before, after := m.findProblematicPages(rules)
	return before == -1 && after == -1
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	rules, manuals := s.parse(lines)
	manuals = slices.Filter(manuals, func(m manual) bool { return m.isCorrect(rules) })
	middlePages := slices.Map(manuals, slices.Middle)
	sum := slices.Sum(middlePages)
	return fmt.Sprintf("%d", sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	rules, manuals := s.parse(lines)
	manuals = slices.Filter(manuals, func(m manual) bool { return !m.isCorrect(rules) })
	slices.ForEach(manuals, func(m manual) { m.fixOrder(rules) })
	middlePages := slices.Map(manuals, slices.Middle)
	sum := slices.Sum(middlePages)
	return fmt.Sprintf("%d", sum)
}
