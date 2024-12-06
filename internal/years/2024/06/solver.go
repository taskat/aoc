package day06

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/slices"
)

// day is the day of the solver
const day = 06

// init registers the solver for day 06
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 06
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) (lab, guard) {
	lab := make(lab, 0, len(lines))
	for _, line := range lines {
		lab = append(lab, slices.Map([]rune(line), parseField))
	}
	startPos := coordinate{x: -1, y: -1}
	for i, row := range lab {
		for j, field := range row {
			if field == guardStart {
				startPos = coordinate{x: j, y: i}
				break
			}
		}
		if startPos.x != -1 {
			break
		}
	}
	guard := guard{
		position: startPos,
		facing:   up,
	}
	return lab, guard
}

type field rune

const (
	empty       field = '.'
	obstruction field = '#'
	visited     field = 'X'
	guardStart  field = '^'
)

func parseField(r rune) field {
	switch r {
	case '.':
		return empty
	case '#':
		return obstruction
	case 'X':
		return visited
	case '^':
		return guardStart
	}
	panic("invalid field")
}

type lab [][]field

func (l lab) isEmpty(c coordinate) bool {
	return l[c.y][c.x] != obstruction
}

func (l lab) isOutOfBounds(c coordinate) bool {
	return c.x < 0 || c.x >= len(l[0]) || c.y < 0 || c.y >= len(l)
}

func (l lab) visit(c coordinate) {
	l[c.y][c.x] = visited
}

type guard struct {
	position coordinate
	facing   direction
}

func (g *guard) move(l lab) bool {
	l.visit(g.position)
	next := g.position.move_new(g.facing)
	if l.isOutOfBounds(next) {
		return false
	}
	if l.isEmpty(next) {
		g.position = next
	} else {
		g.facing = (g.facing + 1) % 4
	}
	return true
}

type coordinate struct {
	x, y int
}

func (c coordinate) move_new(d direction) coordinate {
	switch d {
	case up:
		return coordinate{x: c.x, y: c.y - 1}
	case right:
		return coordinate{x: c.x + 1, y: c.y}
	case down:
		return coordinate{x: c.x, y: c.y + 1}
	case left:
		return coordinate{x: c.x - 1, y: c.y}
	}
	panic("invalid direction")
}

type direction int

const (
	up direction = iota
	right
	down
	left
)

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	lab, guard := s.parse(lines)
	for guard.move(lab) {
	}
	visitedCount := 0
	for _, row := range lab {
		visitedCount += slices.Count(row, func(f field) bool { return f == visited })
	}
	return fmt.Sprint(visitedCount)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
