package day10

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/maps"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 10

// init registers the solver for day 10
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 10
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// mountains represents the mountain height map
type mountains [][]int

// parseMountains parses the input lines into a mountains structure
func parseMountains(lines []string) mountains {
	m := make(mountains, len(lines))
	for i, line := range lines {
		m[i] = slices.Map([]rune(line), stringutils.RuneToInt)
	}
	return m
}

// get returns the height at the given coordinate
func (m mountains) get(c coordinate.Coordinate2D[int]) int {
	return m[c.I()][c.J()]
}

// getTrails returns the number of 9 height points reachable from the start
func (m mountains) getTrails(start coordinate.Coordinate2D[int]) int {
	currents := map[coordinate.Coordinate2D[int]]struct{}{
		start: {},
	}
	currentHeight := m.get(start)
	for currentHeight < 9 {
		neighbors := m.getNextSteps(maps.Keys(currents))
		if len(neighbors) == 0 {
			return 0
		}
		currents = slices.ToMap(neighbors, slices.Repeat(struct{}{}, len(neighbors)))
		currentHeight++
	}
	return len(currents)
}

// getDistinctTrails returns the number of distinct trails to height 9 from the start
func (m mountains) getDistinctTrails(start coordinate.Coordinate2D[int]) int {
	currents := []coordinate.Coordinate2D[int]{start}
	currentHeight := m.get(start)
	for currentHeight < 9 {
		neighbors := m.getNextSteps(currents)
		if len(neighbors) == 0 {
			return 0
		}
		currents = neighbors
		currentHeight++
	}
	return len(currents)
}

// getNeighbors returns the neighbors of the coordinate in the 4 directions
func (m mountains) getNeighbors(c coordinate.Coordinate2D[int]) []coordinate.Coordinate2D[int] {
	directions := []coordinate.Direction{coordinate.Up(), coordinate.Right(),
		coordinate.Down(), coordinate.Left()}
	neighbors := slices.Map(directions, c.Go)
	inBounds := func(c coordinate.Coordinate2D[int]) bool { return c.In2DSlice(len(m[0]), len(m)) }
	neighbors = slices.Filter(neighbors, inBounds)
	return neighbors
}

// getNextSteps returns the neighboring coordinates that are one height higher than any of the currents
func (m mountains) getNextSteps(currents []coordinate.Coordinate2D[int]) []coordinate.Coordinate2D[int] {
	neighbors := make([]coordinate.Coordinate2D[int], 0)
	for _, c := range currents {
		neighbors = append(neighbors, m.getNextFrom(c)...)
	}
	return neighbors
}

// getNextFrom returns the neighboring coordinates that are one height higher than the current
func (m mountains) getNextFrom(from coordinate.Coordinate2D[int]) []coordinate.Coordinate2D[int] {
	neighbors := make([]coordinate.Coordinate2D[int], 0)
	nextHeight := m.get(from) + 1
	for _, n := range m.getNeighbors(from) {
		if m.get(n) == nextHeight {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

// trailheads returns the coordinates of all trailheads (height 0)
func (m mountains) trailheads() []coordinate.Coordinate2D[int] {
	heads := make([]coordinate.Coordinate2D[int], 0)
	for i, row := range m {
		for j, cell := range row {
			if cell == 0 {
				heads = append(heads, coordinate.FromIndexes(i, j))
			}
		}
	}
	return heads
}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) mountains {
	return parseMountains(lines)
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	m := s.parse(lines)
	heads := m.trailheads()
	hikingGoals := slices.Map(heads, m.getTrails)
	sum := slices.Sum(hikingGoals)
	return fmt.Sprint(sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	m := s.parse(lines)
	heads := m.trailheads()
	hikingGoals := slices.Map(heads, m.getDistinctTrails)
	sum := slices.Sum(hikingGoals)
	return fmt.Sprint(sum)
}
