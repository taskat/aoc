package day04

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/slices"
)

// day is the day of the solver
const day = 4

// init registers the solver for day 04
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 04
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) map[rune][]coordinate {
	coordinates := make(map[rune][]coordinate)
	for i, line := range lines {
		for j, c := range line {
			coordinates[c] = append(coordinates[c], coordinate{x: j, y: i})
		}
	}
	return coordinates
}

type coordinate struct {
	x int
	y int
}

func (c coordinate) move_new(d Direction) coordinate {
	c1 := c
	c1.move(d)
	return c1
}

func (c *coordinate) move(d Direction) {
	unit := d.unit()
	c.x += unit.x
	c.y += unit.y
}

type Direction int

const (
	Up Direction = iota
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)

func (d Direction) unit() coordinate {
	switch d {
	case Up:
		return coordinate{0, -1}
	case UpRight:
		return coordinate{1, -1}
	case Right:
		return coordinate{1, 0}
	case DownRight:
		return coordinate{1, 1}
	case Down:
		return coordinate{0, 1}
	case DownLeft:
		return coordinate{-1, 1}
	case Left:
		return coordinate{-1, 0}
	case UpLeft:
		return coordinate{-1, -1}
	}
	panic("invalid direction")
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	coordinates := s.parse(lines)
	count := 0
	for _, c := range coordinates['X'] {
		count += countwords(coordinates, c, "XMAS")
	}
	return fmt.Sprint(count)
}

func countwords(coordinates map[rune][]coordinate, start coordinate, word string) int {
	directions := []Direction{Up, UpRight, Right, DownRight, Down, DownLeft, Left, UpLeft}
	count := 0
	for _, d := range directions {
		if containsWordInDirection(coordinates, start, word, d) {
			count++
		}
	}
	return count
}

func containsWordInDirection(coordinates map[rune][]coordinate, start coordinate, word string, d Direction) bool {
	c := start
	for _, letter := range word {
		if !slices.Contains(coordinates[letter], c) {
			return false
		}
		c.move(d)
	}
	return true
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	coordinates := s.parse(lines)
	count := 0
	for _, c := range coordinates['A'] {
		newcount := countXs(coordinates, c)
		count += newcount
	}
	return fmt.Sprint(count)
}

func countXs(coordinates map[rune][]coordinate, start coordinate) int {
	line1 := []coordinate{start.move_new(UpLeft), start.move_new(DownRight)}
	line2 := []coordinate{start.move_new(UpRight), start.move_new(DownLeft)}
	lines := [][]coordinate{line1, line2}
	for _, line := range lines {
		if !(isPartOfX(coordinates, line[0]) && isPartOfX(coordinates, line[1]) && !isSame(coordinates, line[0], line[1])) {
			return 0
		}
	}
	return 1
}

func isPartOfX(coordinates map[rune][]coordinate, c coordinate) bool {
	return slices.Contains(coordinates['M'], c) || slices.Contains(coordinates['S'], c)
}

func isSame(coordinates map[rune][]coordinate, c1, c2 coordinate) bool {
	return (slices.Contains(coordinates['M'], c1) && slices.Contains(coordinates['M'], c2)) || (slices.Contains(coordinates['S'], c1) && slices.Contains(coordinates['S'], c2))
}
