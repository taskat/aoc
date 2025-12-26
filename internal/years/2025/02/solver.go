package day02

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/containers/set"
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
	return Range{first: first, last: last}
}

// firstPossibleInvalidId returns the first possible invalid ID in the range
func (r Range) firstPossibleInvalidId() int {
	firstStr := fmt.Sprintf("%d", r.first)
	half := firstStr[:len(firstStr)/2]
	return stringutils.Atoi(fmt.Sprintf("%s%s", half, half))
}

// getIDLengthRange returns a new range with the length of IDs
func (r Range) getIDLengthRange() Range {
	first := len(fmt.Sprintf("%d", r.first))
	last := len(fmt.Sprintf("%d", r.last))
	if first == 1 {
		first = 2
	}
	return Range{first: first, last: last}
}

// getRepeatingBlockLengths returns the lengths of repeating blocks of the ID,
// if the ID is in range
func (r Range) getRepeatingBlockLengths() []int {
	lengthSet := set.New[int]()
	idLengthRange := r.getIDLengthRange()
	for length := idLengthRange.first; length <= idLengthRange.last; length++ {
		for blockLength := 1; blockLength <= length/2; blockLength++ {
			if length%(blockLength) == 0 {
				lengthSet.Add(blockLength)
			}
		}
	}
	return lengthSet.ToSlice()
}

// nextPossibleInvalidId generates the next possible invalid ID from the given ID
func (r Range) nextPossibleInvalidId(ID int) int {
	IDStr := fmt.Sprintf("%d", ID)
	half := IDStr[:len(IDStr)/2]
	halfInt := stringutils.Atoi(half)
	halfInt++
	return stringutils.Atoi(fmt.Sprintf("%d%d", halfInt, halfInt))
}

// getDoubleRepetitionIds generates possible invalid IDs that have double repetitions within the range
func (r Range) getDoubleRepetitionIds() []int {
	invalidIds := []int{}
	firstStr := fmt.Sprintf("%d", r.first)
	if len(firstStr)%2 == 1 {
		r.first = stringutils.Atoi(fmt.Sprintf("1%s", strings.Repeat("0", len(firstStr))))
	}
	for possibleId := r.firstPossibleInvalidId(); possibleId <= r.last; possibleId = r.nextPossibleInvalidId(possibleId) {
		if possibleId >= r.first {
			invalidIds = append(invalidIds, possibleId)
		}
	}
	return invalidIds
}

// getRepetitionIds generates possible invalid IDs that have any repetition within the range
func (r Range) getRepetitionIds() []int {
	lengthRange := r.getIDLengthRange()
	blockLengths := r.getRepeatingBlockLengths()
	firstStr := fmt.Sprintf("%d", r.first)
	invalidIds := set.New[int]()
	for _, blockLength := range blockLengths {
		for idLength := lengthRange.first; idLength <= lengthRange.last; idLength++ {
			if idLength%(blockLength) != 0 || idLength <= blockLength {
				continue
			}
			repeatCount := idLength / blockLength
			first := func() (int, int) {
				baseBlock := ""
				if idLength == len(firstStr) {
					baseBlock = firstStr[:blockLength]
				} else {
					baseBlock = fmt.Sprintf("1%s", strings.Repeat("0", blockLength-1))
				}
				id := strings.Repeat(baseBlock, repeatCount)
				return stringutils.Atoi(baseBlock), stringutils.Atoi(id)
			}
			next := func(block int) (int, int) {
				block++
				id := strings.Repeat(fmt.Sprintf("%d", block), repeatCount)
				return block, stringutils.Atoi(id)
			}
			for block, id := first(); id <= r.last; block, id = next(block) {
				if id >= r.first {
					invalidIds.Add(id)
				}
			}
		}
	}
	return invalidIds.ToSlice()
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
	ranges := s.parse(lines)
	invalidIds := []int{}
	for _, r := range ranges {
		invalidIds = append(invalidIds, r.getRepetitionIds()...)
	}
	sum := slices.Sum(invalidIds)
	return fmt.Sprintf("%d", sum)
}
