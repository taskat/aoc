package day10

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 10

// init registers the solver for day 10
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 10
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

type mountains [][]int

func parseMountains(lines []string) mountains {
	m := make(mountains, len(lines))
	for i, line := range lines {
		m[i] = slices.Map([]rune(line), stringutils.RuneToInt)
	}
	return m
}

func (m mountains) get(c coordinate) int {
	return m[c.y][c.x]
}

func (m mountains) getHikingGoals(start coordinate) int {
	currents := map[coordinate]struct{}{
		start: {},
	}
	currentHeight := m.get(start)
	for currentHeight < 9 {
		neighbors := make(map[coordinate]struct{}, 0)
		for c := range currents {
			for _, n := range c.neighbors() {
				if m.inBounds(n) && m.get(n) == currentHeight+1 {
					neighbors[n] = struct{}{}
				}
			}
		}
		if len(neighbors) == 0 {
			return 0
		}
		currents = neighbors
		currentHeight++
	}
	return len(currents)
}

func (m mountains) getDistinctTrails(start coordinate) int {
	currents := []coordinate{start}
	currentHeight := m.get(start)
	for currentHeight < 9 {
		neighbors := make([]coordinate, 0)
		for _, c := range currents {
			for _, n := range c.neighbors() {
				if m.inBounds(n) && m.get(n) == currentHeight+1 {
					neighbors = append(neighbors, n)
				}
			}
		}
		if len(neighbors) == 0 {
			return 0
		}
		currents = neighbors
		currentHeight++
	}
	return len(currents)
}

func (m mountains) inBounds(c coordinate) bool {
	return c.y >= 0 && c.y < len(m) && c.x >= 0 && c.x < len(m[0])
}

func (m mountains) trailheads() []coordinate {
	heads := make([]coordinate, 0)
	for y, row := range m {
		for x, cell := range row {
			if cell == 0 {
				heads = append(heads, coordinate{x, y})
			}
		}
	}
	return heads
}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) mountains {
	return parseMountains(lines)
}

type coordinate struct {
	x, y int
}

func (c coordinate) neighbors() []coordinate {
	return []coordinate{
		{c.x - 1, c.y},
		{c.x + 1, c.y},
		{c.x, c.y - 1},
		{c.x, c.y + 1},
	}
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	m := s.parse(lines)
	heads := m.trailheads()
	hikingGoals := slices.Map(heads, m.getHikingGoals)
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
