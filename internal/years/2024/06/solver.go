package day06

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
	coord "github.com/taskat/aoc/pkg/utils/types/coordinate"
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
func (s *Solver) parse(lines []string) (lab, guard) {
	lab := make(lab, 0, len(lines))
	for _, line := range lines {
		lab = append(lab, slices.Map([]rune(line), parseField))
	}
	startPos := coord.Find2DSliceIndex(lab, guardStart.predicatorFactory())
	guard := coord.NewActor(startPos, coord.Up())
	return lab, *guard
}

// field represents a cell in the lab
type field rune

const (
	empty       field = '.'
	obstruction field = '#'
	visited     field = 'X'
	guardStart  field = '^'
)

// predicatorFactory creates a predicate function that returns true if the cell
// is the requested field
func (f field) predicatorFactory() func(field) bool {
	return func(cell field) bool {
		return cell == f
	}
}

// parseField converts a rune to a field
// The '.', '#', 'X', and '^' runes are valid fields
// Any other rune will cause a panic
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

// lab represents the lab layout, it is a 2D slice of fields
type lab [][]field

// isEmpty checks if the cell at the given coordinate is empty
func (l lab) isEmpty(c coord.Coordinate2D[int]) bool {
	return l[c.I()][c.J()] != obstruction
}

// isOutOfBounds checks if the coordinate is out of bounds of the lab
func (l lab) isOutOfBounds(c coord.Coordinate2D[int]) bool {
	return !c.In2DSlice(len(l[0]), len(l))
}

// leadsToCircle checks if starting from the given coordinate and direction
// will eventually lead to a circle (i.e., revisiting the same cell facing the same direction)
func (l lab) leadsToCircle(g guard) bool {
	visitedFields := make(map[coord.Coordinate2D[int]][]coord.Direction)
	visitedFields[g.Position] = []coord.Direction{g.Direction}
	for l.move(&g, false) {
		dir, ok := visitedFields[g.Position]
		if ok && slices.Contains(dir, g.Direction) {
			return true
		}
		visitedFields[g.Position] = append(dir, g.Direction)
	}
	return false
}

// move moves the guard in the lab according to the rules
func (l lab) move(g *guard, visit bool) bool {
	if visit {
		l.visit(g.Position)
	}
	next := g.Next()
	if l.isOutOfBounds(next) {
		return false
	}
	if l.isEmpty(next) {
		g.Position = next
	} else {
		g.Direction = g.Direction.TurnRight()
	}
	return true
}

// String implements the fmt.Stringer interface
func (l lab) String() string {
	result := ""
	for _, row := range l {
		for _, f := range row {
			result += string(f)
		}
		result += "\n"
	}
	return result
}

// visit marks the cell at the given coordinate as visited
func (l lab) visit(c coord.Coordinate2D[int]) {
	l[c.I()][c.J()] = visited
}

type guard = coord.Actor[int]

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	lab, guard := s.parse(lines)
	for lab.move(&guard, true) {
	}
	visitedCount := coordinate.CountIn2DSlice(lab, visited.predicatorFactory())
	return fmt.Sprint(visitedCount)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	lab, guard := s.parse(lines)
	count := 0
	for {
		next := guard.Next()
		if lab.isOutOfBounds(next) {
			break
		}
		if lab[next.I()][next.J()] != empty {
			lab.move(&guard, true)
			continue
		}
		oldField := lab[next.I()][next.J()]
		lab[next.I()][next.J()] = obstruction
		if lab.leadsToCircle(guard) {
			count++
		}
		lab[next.I()][next.J()] = oldField
		if !lab.move(&guard, true) {
			break
		}
	}
	return fmt.Sprint(count)
}
