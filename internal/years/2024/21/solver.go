package day21

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/intutils"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 21

// init registers the solver for day 21
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 21
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []code {
	return slices.Map(lines, newCode)
}

type code []rune

func newCode(s string) code {
	return code(s)
}

func (c code) buttons(pads []keypad) map[string]int {
	codeblocks := make(map[string]int)
	codeblocks[c.String()] = 1
	for i := 0; i < len(pads); i++ {
		nextCodeBlocks := make(map[string]int)
		for codeBlock, count := range codeblocks {
			newCodeBlocks := pads[i].typeCode(newCode(codeBlock))
			for newCodeBlock, newCount := range newCodeBlocks {
				nextCodeBlocks[newCodeBlock] += (newCount * count)
			}
		}
		codeblocks = nextCodeBlocks
	}
	return codeblocks
}

func (c code) complexity(pads []keypad) int {
	finalCode := c.buttons(pads)
	numPart := strings.ReplaceAll(string(c), "A", "")
	numValue := stringutils.Atoi(numPart)
	length := 0
	for codeBlock, count := range finalCode {
		fmt.Println(codeBlock, count)
		length += len(codeBlock) * count
	}
	return length * numValue
}

func (c code) String() string {
	return string(c)
}

type coord = coordinate.Coordinate2D[int]

type dir = coordinate.Direction

type keypad map[rune]coord

func numPad() keypad {
	return map[rune]coord{
		'7': {X: 0, Y: 0},
		'8': {X: 1, Y: 0},
		'9': {X: 2, Y: 0},
		'4': {X: 0, Y: 1},
		'5': {X: 1, Y: 1},
		'6': {X: 2, Y: 1},
		'1': {X: 0, Y: 2},
		'2': {X: 1, Y: 2},
		'3': {X: 2, Y: 2},
		'0': {X: 1, Y: 3},
		'A': {X: 2, Y: 3},
	}
}

func directionalPad() keypad {
	return map[rune]coord{
		'^': {X: 1, Y: 0},
		'A': {X: 2, Y: 0},
		'<': {X: 0, Y: 1},
		'v': {X: 1, Y: 1},
		'>': {X: 2, Y: 1},
	}
}

func (k keypad) buttons(from, to rune) code {
	fromCoord := k[from]
	toCoord := k[to]
	dist := toCoord.Sub(fromCoord)
	buttons := make(code, 0, intutils.Abs(dist.X)+intutils.Abs(dist.Y)+1)
	appendN := func(n int, r rune) {
		for i := 0; i < n; i++ {
			buttons = append(buttons, r)
		}
	}
	if (to == '1' || to == '4' || to == '7') && (from == '0' || from == 'A') {
		appendN(-dist.Y, '^')
		appendN(-dist.X, '<')
	} else if (from == '^' || from == 'A') && to == '<' {
		appendN(dist.Y, 'v')
		appendN(-dist.X, '<')
	} else {
		if dist.X < 0 {
			appendN(-dist.X, '<')
		}
		if (from == '1' || from == '4' || from == '7') && (to == '0' || to == 'A') {
			appendN(dist.X, '>')
			appendN(dist.Y, 'v')
		} else if from == '<' && (to == '^' || to == 'A') {
			appendN(dist.X, '>')
			appendN(-dist.Y, '^')
		} else {
			if dist.Y > 0 {
				appendN(dist.Y, 'v')
			}
			if dist.Y < 0 {
				appendN(-dist.Y, '^')
			}
			if dist.X > 0 {
				appendN(dist.X, '>')
			}
		}
	}
	buttons = append(buttons, 'A')
	return buttons
}

func (k keypad) typeCode(code code) map[string]int {
	codeBlocks := make(map[string]int, 0)
	firstBlock := k.buttons('A', code[0])
	codeBlocks[firstBlock.String()] = 1
	for i := 0; i < len(code)-1; i++ {
		newBlock := k.buttons(code[i], code[i+1])
		codeBlocks[newBlock.String()]++
	}
	return codeBlocks
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	codes := s.parse(lines)
	pads := []keypad{numPad(), directionalPad(), directionalPad()}
	complexities := slices.Map(codes, func(c code) int { return c.complexity(pads) })
	return fmt.Sprint(slices.Sum(complexities))
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	codes := s.parse(lines)
	pads := make([]keypad, 26)
	pads[0] = numPad()
	for i := 0; i < 25; i++ {
		pads[i+1] = directionalPad()
	}
	complexities := slices.Map(codes, func(c code) int { return c.complexity(pads) })
	return fmt.Sprint(slices.Sum(complexities))
}
