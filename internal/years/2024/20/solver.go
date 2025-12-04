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

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) (*graph.Graph[coordinate.Coordinate2D[int]], *graph.BaseNode[coordinate.Coordinate2D[int]], coordinate.Coordinate2D[int]) {
	nodes := set.New[coordinate.Coordinate2D[int]]()
	var start, end coordinate.Coordinate2D[int]
	for i, line := range lines {
		for j, cell := range line {
			if isWall(cell) {
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
	g := graph.FromSliceOfCoordinates(nodes.ToSlice(), coordinate.Straights())
	g.AddNode(startNode)
	return g, startNode, end
}

// isWall checks if a cell is a wall, represented by '#'
func isWall(cell rune) bool {
	return cell == '#'
}

// getNeighbors gets the walkable neighbors of a coordinate in the graph
func getNeighbors(g *graph.Graph[coordinate.Coordinate2D[int]], c coordinate.Coordinate2D[int]) []coordinate.Coordinate2D[int] {
	neighbors := maps.Values(c.Neighbors(coordinate.Straights()))
	neighbors = slices.Filter(neighbors, func(c coordinate.Coordinate2D[int]) bool { return g.HasNode(c) })
	return neighbors
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	g, startNode, _ := s.parse(lines)
	distances := g.Distances(startNode)
	saves := map[int]int{}
	for i := range lines {
		for j := range lines[i] {
			if !isWall(rune(lines[i][j])) {
				continue
			}
			c := coordinate.FromIndexes(i, j)
			neighbors := getNeighbors(g, c)
			if len(neighbors) != 2 {
				continue
			}
			neighborDistances := slices.Map(neighbors, func(c coordinate.Coordinate2D[int]) int { return distances[c] })
			savedTime := intutils.Abs(neighborDistances[0]-neighborDistances[1]) - 2
			saves[savedTime]++
		}
	}
	saves = maps.Filter(saves, func(saved, _ int) bool { return saved >= 100 })
	return fmt.Sprint(maps.Sum(saves))
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	g, startNode, _ := s.parse(lines)
	distances := g.Distances(startNode)
	saves := map[int]int{}
	nodes := maps.Keys(g.GetNodes())
	for i := range nodes {
		for j := i + 1; j < len(nodes); j++ {
			coord1 := nodes[i]
			coord2 := nodes[j]
			manhattanDistance := coord1.ManhattanDistance(coord2)
			if manhattanDistance > 20 {
				continue
			}
			savedTime := intutils.Abs(distances[coord1]-distances[coord2]) - manhattanDistance
			saves[savedTime]++
		}
	}
	saves = maps.Filter(saves, func(saved, _ int) bool { return saved >= 100 })
	return fmt.Sprint(maps.Sum(saves))
}
