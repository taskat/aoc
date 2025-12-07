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
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 21
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []code {
	return slices.Map(lines, newCode)
}

// code represents a sequence of buttons pressed
type code []rune

// newCode creates a new code from a string
func newCode(s string) code {
	return code(s)
}

// addHorizontal appends n horizontal moves to the code
func (c *code) addHorizontal(n int) {
	if n < 0 {
		c.appendN(-n, '<')
	} else {
		c.appendN(n, '>')
	}
}

// addVertical appends n vertical moves to the code
func (c *code) addVertical(n int) {
	if n < 0 {
		c.appendN(-n, 'v')
	} else {
		c.appendN(n, '^')
	}
}

// appendN appends n times the rune r to the code
func (c *code) appendN(n int, r rune) {
	*c = append(*c, slices.Repeat(r, n)...)
}

// buttons returns the mapping of button sequences for a given code
// on the provided keypads. It contains which button sequences are used and how many times
func (c code) buttons(pads []keypad) map[string]int {
	codeblocks := make(map[string]int)
	codeblocks[c.String()] = 1
	for i := range pads {
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

// complexity calculates the complexity of the code given the keypads
// It is the product of the length of the final code and the numeric value of the code
func (c code) complexity(pads []keypad) int {
	finalCode := c.buttons(pads)
	numPart := strings.ReplaceAll(string(c), "A", "")
	numValue := stringutils.Atoi(numPart)
	length := 0
	for codeBlock, count := range finalCode {
		length += len(codeBlock) * count
	}
	return length * numValue
}

// String returns the string representation of the code
func (c code) String() string {
	return string(c)
}

// keypad maps buttons to their coordinates
type keypad struct {
	keys    map[rune]coordinate.Coordinate2D[int]
	numeric bool
}

// numPad returns the numeric keypad layout
func numPad() keypad {
	return keypad{
		keys: map[rune]coordinate.Coordinate2D[int]{
			'7': {X: 0, Y: 0},
			'8': {X: 1, Y: 0},
			'9': {X: 2, Y: 0},
			'4': {X: 0, Y: -1},
			'5': {X: 1, Y: -1},
			'6': {X: 2, Y: -1},
			'1': {X: 0, Y: -2},
			'2': {X: 1, Y: -2},
			'3': {X: 2, Y: -2},
			'0': {X: 1, Y: -3},
			'A': {X: 2, Y: -3},
		},
		numeric: true,
	}
}

// directionalPad returns the directional keypad layout
func directionalPad() keypad {
	return keypad{
		keys: map[rune]coordinate.Coordinate2D[int]{
			'^': {X: 1, Y: 0},
			'A': {X: 2, Y: 0},
			'<': {X: 0, Y: -1},
			'v': {X: 1, Y: -1},
			'>': {X: 2, Y: -1},
		},
		numeric: false,
	}
}

// buttons returns the sequence of button presses necessary to guide the robot
// from one button to another
func (k keypad) buttons(from, to rune) code {
	fromCoord := k.keys[from]
	toCoord := k.keys[to]
	dist := toCoord.Sub(fromCoord)
	buttons := make(code, 0, intutils.Abs(dist.X)+intutils.Abs(dist.Y)+1)
	if k.isVerticalFirst(from, to) {
		buttons.addVertical(dist.Y)
		buttons.addHorizontal(dist.X)
	} else if k.isHorizontalFirst(from, to) {
		buttons.addHorizontal(dist.X)
		buttons.addVertical(dist.Y)
	} else if dist.X < 0 {
		buttons.addHorizontal(dist.X)
		buttons.addVertical(dist.Y)
	} else {
		buttons.addVertical(dist.Y)
		buttons.addHorizontal(dist.X)
	}
	buttons = append(buttons, 'A')
	return buttons
}

// isVerticalFirst determines if vertical moves should be made before horizontal ones
// based on the keypad layout and the from/to buttons. This is used to minimize the number of moves.
// If the keypad is numeric, vertical moves are prioritized when moving to the left column (1,4,7)
// from the bottom row (0,A). If the keypad is directional, vertical moves are prioritized when to the
// left column (<) from the top row (^,A).
func (k keypad) isVerticalFirst(from, to rune) bool {
	if k.numeric {
		return (from == '0' || from == 'A') && (to == '1' || to == '4' || to == '7')
	}
	return (from == '^' || from == 'A') && (to == '<')
}

// isHorizontalFirst determines if horizontal moves should be made before vertical ones
// based on the keypad layout and the from/to buttons. This is used to minimize the number of moves.
// If the keypad is numeric, horizontal moves are prioritized when moving from the left column (1,4,7)
// to the bottom row (0,A). If the keypad is directional, horizontal moves are prioritized when moving
// from the left column (<) to the top row (^,A).
func (k keypad) isHorizontalFirst(from, to rune) bool {
	if k.numeric {
		return (from == '1' || from == '4' || from == '7') && (to == '0' || to == 'A')
	}
	return (from == '<') && (to == '^' || to == 'A')
}

// typeCode returns the mapping of button sequences for a given code
// on this keypad. It contains which button sequences are used and how many times
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
