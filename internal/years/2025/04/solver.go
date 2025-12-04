package day04

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/containers/set"
	"github.com/taskat/aoc/pkg/utils/maps"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 4

// init registers the solver for day 04
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 04
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) grid {
	g := grid{
		coordinates: set.New[coordinate.Coordinate2D[int]](),
		width:       len(lines[0]),
		height:      len(lines),
	}
	for i, line := range lines {
		for j, char := range line {
			if char == '@' {
				g.coordinates.Add(coordinate.FromIndexes(i, j))
			}
		}
	}
	return g
}

// grid contains the coordinates of the paper rolls, and the size of the grid
type grid struct {
	coordinates set.Set[coordinate.Coordinate2D[int]]
	width       int
	height      int
}

// accessible checks if a coordinate is accessible, which means it has less than 4 neighbors
func (g *grid) accessible(c coordinate.Coordinate2D[int]) bool {
	neighbors := c.Neighbors(coordinate.Directions())
	neighbors = maps.Filter(neighbors, func(_ coordinate.Direction, c coordinate.Coordinate2D[int]) bool {
		return c.In2DSlice(g.width, g.height)
	})
	for d, n := range neighbors {
		if !g.coordinates.Contains(n) {
			delete(neighbors, d)
		}
	}
	return len(neighbors) < 4
}

// countAccessible counts the number of accessible coordinates in the grid
func (g *grid) countAccessible() int {
	coords := g.coordinates.ToSlice()
	coords = slices.Filter(coords, g.accessible)
	return len(coords)
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	grid := s.parse(lines)
	return fmt.Sprintf("%d", grid.countAccessible())
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
