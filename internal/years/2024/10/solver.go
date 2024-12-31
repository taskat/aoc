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

// coord is an alias for coordinate.Coordinate2D[int]
type coord = coordinate.Coordinate2D[int]

type mountains [][]int

func parseMountains(lines []string) mountains {
	m := make(mountains, len(lines))
	for i, line := range lines {
		m[i] = slices.Map([]rune(line), stringutils.RuneToInt)
	}
	return m
}

func (m mountains) get(c coord) int {
	return m[c.Y][c.X]
}

func (m mountains) getTrails(start coord) int {
	currents := map[coord]struct{}{
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

func (m mountains) getDistinctTrails(start coord) int {
	currents := []coord{start}
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
func (m mountains) getNeighbors(c coord) []coord {
	directions := []coordinate.Direction{coordinate.Up(), coordinate.Right(),
		coordinate.Down(), coordinate.Left()}
	neighbors := slices.Map(directions, c.Go)
	inBounds := func(c coord) bool { return coord.In2DSlice(c, len(m[0]), len(m)) }
	neighbors = slices.Filter(neighbors, inBounds)
	return neighbors
}

func (m mountains) getNextSteps(currents []coord) []coord {
	neighbors := make([]coord, 0)
	for _, c := range currents {
		neighbors = append(neighbors, m.getNextFrom(c)...)
	}
	return neighbors
}

func (m mountains) getNextFrom(from coord) []coord {
	neighbors := make([]coord, 0)
	nextHeight := m.get(from) + 1
	for _, n := range m.getNeighbors(from) {
		if m.get(n) == nextHeight {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func (m mountains) trailheads() []coord {
	heads := make([]coord, 0)
	for y, row := range m {
		for x, cell := range row {
			if cell == 0 {
				heads = append(heads, coordinate.NewCoordinate2D(x, y))
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
