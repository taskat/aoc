package day20

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/containers/set"
	"github.com/taskat/aoc/pkg/utils/graph"
	"github.com/taskat/aoc/pkg/utils/intutils"
	"github.com/taskat/aoc/pkg/utils/maps"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 20

// init registers the solver for day 20
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 20
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

type coord = coordinate.Coordinate2D[int]

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) (*graph.Graph[coord], *graph.BaseNode[coord], coord) {
	nodes := set.New[coord]()
	var start, end coord
	for i, line := range lines {
		for j, cell := range line {
			if cell == '#' {
				continue
			}
			c := coordinate.FromIndexes(i, j)
			switch cell {
			case 'S':
				start = c
			case 'E':
				end = c
			}
			nodes.Add(c)
		}
	}
	startNode := graph.NewBaseNode(start)
	g := graph.New[coord]()
	g.AddNode(startNode)
	for _, c := range nodes.ToSlice() {
		possibleNeighbors := c.Neighbors(coordinate.Straights())
		neighbors := maps.Filter(possibleNeighbors, func(dir coordinate.Direction, c coord) bool { return nodes.Contains(c) })
		for _, neighbor := range neighbors {
			g.AddNode(graph.NewBaseNode(c))
			g.AddNode(graph.NewBaseNode(neighbor))
			g.AddEdge(c, neighbor, 1)
		}
	}
	return g, startNode, end
}

func hasOpposites(dirs []coordinate.Direction) bool {
	if len(dirs) < 2 {
		return false
	}
	for i := 0; i < len(dirs); i++ {
		for j := i + 1; j < len(dirs); j++ {
			if dirs[i].Opposite() == dirs[j] {
				return true
			}
		}
	}
	return false
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	g, startNode, _ := s.parse(lines)
	distances := g.Distances(startNode)
	saves := map[int]int{}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != '#' {
				continue
			}
			c := coordinate.FromIndexes(i, j)
			neighbors := maps.Values(c.Neighbors(coordinate.Straights()))
			neighbors = slices.Filter(neighbors, func(c coord) bool { return g.HasNode(c) })
			if len(neighbors) != 2 {
				continue
			}
			neighborDistances := slices.Map(neighbors, func(c coord) int { return distances[c] })
			saved := intutils.Abs(neighborDistances[0]-neighborDistances[1]) - 2
			saves[saved]++
		}
	}
	saves = maps.Filter(saves, func(saved, _ int) bool { return saved >= 100 })
	savedRoutes := maps.Values(saves)
	return fmt.Sprint(slices.Sum(savedRoutes))
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	g, startNode, _ := s.parse(lines)
	distances := g.Distances(startNode)
	saves := map[int]int{}
	nodes := maps.Keys(g.GetNodes())
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			coord1 := nodes[i]
			coord2 := nodes[j]
			manhattanDistance := coord1.ManhattanDistance(coord2)
			if manhattanDistance > 20 {
				continue
			}
			saved := intutils.Abs(distances[coord1]-distances[coord2]) - manhattanDistance
			saves[saved]++
		}
	}
	saves = maps.Filter(saves, func(saved, _ int) bool { return saved >= 100 })
	savedRoutes := maps.Values(saves)
	return fmt.Sprint(slices.Sum(savedRoutes))
}
