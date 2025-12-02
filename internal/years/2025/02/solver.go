package day02

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/iterutils"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 2

// init registers the solver for day 02
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 02
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []Range {
	parts := strings.Split(lines[0], ",")
	return iterutils.ToSlice(iterutils.Map(iterutils.NewFromSlice(parts), parseRange))
}

// Range represents a range of IDs
type Range struct {
	first int
	last  int
}

// parseRange parses a range from a string, e.g. "1-3" -> Range{1, 3}
func parseRange(s string) Range {
	parts := strings.Split(s, "-")
	first := stringutils.Atoi(parts[0])
	last := stringutils.Atoi(parts[1])
	if len(parts[0])%2 == 1 {
		first = stringutils.Atoi(fmt.Sprintf("1%s", strings.Repeat("0", len(parts[0]))))
	}
	return Range{first: first, last: last}
}

// firstPossibleInvalidId returns the first possible invalid ID in the range
func (r Range) firstPossibleInvalidId() int {
	firstStr := fmt.Sprintf("%d", r.first)
	half := firstStr[:len(firstStr)/2]
	return stringutils.Atoi(fmt.Sprintf("%s%s", half, half))
}

// nextPossibleInvalidId generates the next possible invalid ID from the given ID
func (r Range) nextPossibleInvalidId(ID int) int {
	IDStr := fmt.Sprintf("%d", ID)
	half := IDStr[:len(IDStr)/2]
	halfInt := stringutils.Atoi(half)
	halfInt++
	return stringutils.Atoi(fmt.Sprintf("%d%d", halfInt, halfInt))
}

// getDoubleRepetitionIds generates possible invalid IDs from half IDs within the range
func (r Range) getDoubleRepetitionIds() []int {
	invalidIds := []int{}
	for possibleId := r.firstPossibleInvalidId(); possibleId <= r.last; possibleId = r.nextPossibleInvalidId(possibleId) {
		if possibleId >= r.first {
			invalidIds = append(invalidIds, possibleId)
		}
	}
	return invalidIds
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	ranges := s.parse(lines)
	invalidIds := []int{}
	for _, r := range ranges {
		invalidIds = append(invalidIds, r.getDoubleRepetitionIds()...)
	}
	sum := slices.Sum(invalidIds)
	return fmt.Sprintf("%d", sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
